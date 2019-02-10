package main

import (
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
	oauthServer, err := oauthServerFromEnv(loginPath, tokenPath)
	if err != nil {
		log.Fatalf("error setting up oauth server: %s", err)
	}
	var tokenParser authTokenParser = oauthServer

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
		log.Printf("authorized call by %q on behalf of %q", token.ClientID, token.User)
		_, _ = io.Copy(os.Stderr, r.Body) // TODO
		http.Error(rw, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	})

	log.Println("setup successful, starting to listen")
	err = http.ListenAndServe("127.0.0.1:18917", mux)
	log.Println("finished serving with err =", err)
}
