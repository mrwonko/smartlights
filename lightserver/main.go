package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/mrwonko/smartlights/config"
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
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
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

	var tokenParser authTokenParser = oauthServer
	syncChan := func(ctx context.Context, googleAPIKey, user string) chan<- struct{} {
		res := make(chan struct{}, 16)
		go sendSyncRequests(ctx, http.DefaultClient, res, googleAPIKey, user)
		return res
	}(ctx, googleAPIKey, user)

	mux := http.NewServeMux()

	mux.HandleFunc(loginPath, oauthServer.serveLogin)
	mux.HandleFunc(tokenPath, oauthServer.serveToken)
	mux.HandleFunc("/smartlights/fulfillment", func(rw http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			log.Print("fulfillment called without auth")
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(auth, bearerPrefix) {
			log.Print("fulfillment called with non-bearer auth")
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		auth = strings.TrimPrefix(auth, bearerPrefix)
		token, err := tokenParser.parseAuthToken(auth, time.Now(), typeAccessToken, "")
		if err != nil {
			log.Printf("fulfillment called with invalid token: %s", err)
			http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		if r.Method != http.MethodPost {
			log.Printf("fulfillment called with invalid method %q", r.Method)
			http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("fulfillment failed to parse body: %s", err)
			http.Error(rw, "invalid JSON", http.StatusBadRequest)
			return
		}
		if len(req.Inputs) != 1 {
			log.Printf("fulfillment invalid input count: %v", req.Inputs)
			http.Error(rw, "invalid input count", http.StatusUnprocessableEntity)
			return
		}
		input := &req.Inputs[0]
		switch input.Intent {
		case intentSync:
			resp := response{
				RequestID: req.RequestID,
				Payload: responsePayloadSync{
					AgentUserID: token.User,
					Devices:     devices,
				},
			}
			if err := json.NewEncoder(io.MultiWriter(rw, os.Stderr)).Encode(&resp); err != nil {
				log.Printf("fulfillment failed to send response: %s", err)
			}
			log.Printf("fulfillment sync successful")
		case intentDisconnect:
			log.Printf("TODO implement disconnect")
			http.Error(rw, "TODO implement disconnect", http.StatusNotImplemented)
		case intentExecute:
			inputPayload := requestPayloadExecute{}
			if err := json.Unmarshal(input.Payload, &inputPayload); err != nil {
				log.Printf("fulfillment execute json parse: %s", err)
				http.Error(rw, "failed to parse body", http.StatusBadRequest)
				return
			}
			type deviceErr struct {
				device config.ID
				err    error
			}
			deviceErrs, deviceErrsMapChan := func(ctx context.Context) (chan<- deviceErr, <-chan map[config.ID]error) {
				reqChan := make(chan deviceErr)
				resChan := make(chan map[config.ID]error)
				go func(ctx context.Context, reqChan <-chan deviceErr, resChan chan<- map[config.ID]error) {
					defer close(resChan)
					res := map[config.ID]error{}
					for {
						select {
						case req, ok := <-reqChan:
							if !ok {
								resChan <- res
								return
							}
							if res[req.device] == nil {
								res[req.device] = req.err
							}
						case <-ctx.Done():
							return
						}
					}
				}(ctx, reqChan, resChan)
				return reqChan, resChan
			}(r.Context())
			sendErr := func(err errorCode) {
				if err := json.NewEncoder(rw).Encode(response{
					RequestID: req.RequestID,
					Payload: responsePayloadError{
						ErrorCode: err,
					},
				}); err != nil {
					log.Printf("fulfillment execute error response: %s", err)
				}
			}
			var wg sync.WaitGroup
			allDevices := map[config.ID]*config.Light{}
			for i := range inputPayload.Commands {
				command := &inputPayload.Commands[i]

				devices := make(map[config.ID]*config.Light, len(command.Devices))
				for i := range command.Devices {
					var id int
					id, err = strconv.Atoi(command.Devices[i].ID)
					if err != nil {
						log.Printf("fulfillment execute device ID %q parse: %s", command.Devices[i].ID, err)
						sendErr(errorCodeDeviceNotFound)
						return
					}
					light := config.Lights[config.ID(id)]
					if light == nil {
						log.Printf("unknown device %d", id)
						sendErr(errorCodeDeviceNotFound)
						return
					}
					devices[config.ID(id)] = light
					allDevices[config.ID(id)] = light
				}

				ons := make([]bool, 0, len(command.Execution))
				for i := range command.Execution {
					e := &command.Execution[i]
					switch e.Command {
					case "action.devices.commands.OnOff":
						on, ok := e.Params["on"].(bool)
						if !ok {
							log.Printf(`fulfillment execute OnOff command without "on" param`)
							sendErr(errorCodeProtocolError)
							return
						}
						ons = append(ons, on)
					default:
						log.Printf("unsupported command %q", e.Command)
						sendErr(errorCodeFunctionNotSupported)
						return
					}
				}
				if len(ons) > 1 {
					log.Printf("too many ons: %v", ons)
					sendErr(errorCodeProtocolError)
					continue
				}

				for _, on := range ons {
					for id := range devices {
						wg.Add(1)
						go func(id config.ID, on bool) {
							deviceErrs <- deviceErr{
								device: id,
								err:    pc.OnOff(r.Context(), id, on),
							}
							wg.Done()
						}(id, on)
					}
				}
			}
			wg.Wait()
			close(deviceErrs)
			deviceErrsMap := <-deviceErrsMapChan
			payload := responsePayloadExecute{
				Commands: make([]responsePayloadExecuteCommand, len(deviceErrsMap)),
			}
			i := 0
			for id, err := range deviceErrsMap {
				cmd := &payload.Commands[i]
				i++
				cmd.IDs = []string{strconv.Itoa(int(id))}
				if err == nil {
					cmd.Status = statusPending
				} else {
					log.Printf("error toggling device %d: %s", id, err)
					msg := err.Error()
					cmd.DebugString = &msg
					cmd.Status = statusError
					cmd.ErrorCode = errorCodeTransientError
				}
			}
			if err := json.NewEncoder(io.MultiWriter(rw, os.Stderr)).Encode(response{
				RequestID: req.RequestID,
				Payload:   payload,
			}); err != nil {
				log.Printf("fulfillment execute send response: %s", err)
				return
			}
			log.Printf("successful execute for %d devices", len(allDevices))
		case intentQuery:
			log.Printf("TODO implement query")
			http.Error(rw, "TODO implement query", http.StatusNotImplemented)
		default:
			log.Printf("fulfillment unsupported intent %q", input.Intent)
			http.Error(rw, "unsupported intent", http.StatusNotImplemented)
		}
	})

	log.Println("setup successful, starting to listen & requesting sync")
	go func(syncChan chan<- struct{}) {
		time.Sleep(time.Second) // allow server to finish starting
		syncChan <- struct{}{}
	}(syncChan)
	go func() {
		err := pc.ReceiveState(ctx, func(ctx context.Context, msg *config.StateMessage) {
			if err := reportState(ctx, http.DefaultClient, user, googleCreds.privateKey, googleCreds.clientEmail); err != nil {
				log.Printf("error reporting state: %s", err)
			}
		})
		if err != nil {
			log.Printf("fatal error receiving states: %s", err)
		}
	}()
	server := http.Server{
		Addr:    "127.0.0.1:18917",
		Handler: mux,
	}
	serveErrChan := make(chan error)
	go func() {
		serveErrChan <- server.ListenAndServe()
	}()
	<-ctx.Done()
	log.Printf("shutting down server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = server.Shutdown(shutdownCtx); err != nil {
		log.Printf("failed to shut down server: %s", err)
	} else {
		err = <-serveErrChan
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
