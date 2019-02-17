package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

	var tokenParser authTokenParser = oauthServer
	syncChan := func(ctx context.Context, googleAPIKey, user string) chan<- struct{} {
		res := make(chan struct{}, 16)
		go sendSyncRequests(ctx, http.DefaultClient, res, googleAPIKey, user)
		return res
	}(context.TODO(), googleAPIKey, user)

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
			log.Printf("TODO implement execute")
			http.Error(rw, "TODO implement execute", http.StatusNotImplemented)
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
	err = http.ListenAndServe("127.0.0.1:18917", mux)
	log.Println("finished serving with err =", err)
}
