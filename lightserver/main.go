package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrwonko/smartlights/internal/protocol"
	"golang.org/x/sync/errgroup"
)

func main() {
	const (
		loginPath = "/smartlights/login"
		tokenPath = "/smartlights/token"
	)
	googleAPIKey := os.Getenv("GOOGLE_API_KEY")
	if googleAPIKey == "" {
		log.Fatalf("missing GOOGLE_API_KEY")
	}
	googleCreds, err := loadGoogleCredentials()
	if err != nil {
		log.Fatalf("reading Google credentials: %s", err)
	}

	oauthServer, err := oauthServerFromEnv(loginPath, tokenPath)
	if err != nil {
		log.Fatalf("error setting up oauth server: %s", err)
	}
	user := func(users map[string][]byte) string {
		for u := range users {
			return u
		}
		return ""
	}(oauthServer.userPasswordHashes)

	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 2)
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	signal.Notify(sigChan, sigs...)
	go func() {
		<-sigChan
		signal.Reset(sigs...)
		cancel()
	}()

	pc, err := newPubsubClient(ctx)
	if err != nil {
		log.Fatalf("creating pubsub client: %s", err)
	}
	defer func() {
		if err = pc.Close(); err != nil {
			log.Printf("failed to close pubsub client: %s", err)
		}
	}()

	var eg errgroup.Group

	syncChan := func(ctx context.Context, googleAPIKey, user string) chan<- struct{} {
		res := make(chan struct{}, 16)
		eg.Go(func() error {
			return sendSyncRequests(ctx, http.DefaultClient, res, googleAPIKey, user)
		})
		return res
	}(ctx, googleAPIKey, user)

	mux := http.NewServeMux()

	mux.HandleFunc(loginPath, oauthServer.serveLogin)
	mux.HandleFunc(tokenPath, oauthServer.serveToken)
	mux.HandleFunc("/smartlights/fulfillment", serveFulfillment(oauthServer, pc))

	log.Println("setup successful, starting to listen & requesting sync")
	go func(syncChan chan<- struct{}) {
		time.Sleep(time.Second) // allow server to finish starting
		syncChan <- struct{}{}
	}(syncChan)
	eg.Go(func() error {
		tokenCache := newAccessTokenCache(http.DefaultClient, googleCreds.privateKey, googleCreds.clientEmail)
		err := pc.ReceiveState(ctx, func(ctx context.Context, msg *protocol.StateMessage) {
			if err := reportState(ctx, tokenCache, http.DefaultClient, user); err != nil {
				log.Printf("error reporting state: %s", err)
			}
		})
		if err != nil {
			return fmt.Errorf("fatal error receiving states: %w", err)
		}
		return nil
	})
	server := http.Server{
		Addr:    "127.0.0.1:18917",
		Handler: mux,
	}
	eg.Go(func() error {
		return server.ListenAndServe()
	})
	<-ctx.Done()
	log.Printf("shutting down server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = server.Shutdown(shutdownCtx); err != nil {
		log.Printf("failed to shut down server: %s", err)
	} else {
		err = eg.Wait()
		log.Printf("finished serving with err=%s", err)
	}
}

type googleCredentials struct {
	privateKey  *rsa.PrivateKey
	clientEmail string
}

func loadGoogleCredentials() (_ *googleCredentials, resErr error) {
	googleCredsFilename := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if googleCredsFilename == "" {
		return nil, fmt.Errorf("missing GOOGLE_APPLICATION_CREDENTIALS")
	}
	googleCredsFile, err := os.Open(googleCredsFilename)
	if err != nil {
		return nil, fmt.Errorf("opening Google credentials file %q: %w", googleCredsFilename, err)
	}
	defer func() {
		closeErr := googleCredsFile.Close()
		if closeErr != nil && resErr == nil {
			resErr = fmt.Errorf("closing %q: %w", googleCredsFilename, err)
		}
	}()
	var data struct {
		PrivateKey  string `json:"private_key"`
		ClientEmail string `json:"client_email"`
	}
	if err = json.NewDecoder(googleCredsFile).Decode(&data); err != nil {
		return nil, fmt.Errorf("parsing creds file %q: %w", googleCredsFilename, err)
	}
	res := googleCredentials{
		clientEmail: data.ClientEmail,
	}
	block, _ := pem.Decode([]byte(data.PrivateKey))
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing private key: %w", err)
	}
	var ok bool
	res.privateKey, ok = privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key is not an RSA Private Key")
	}
	return &res, nil
}
