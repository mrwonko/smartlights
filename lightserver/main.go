package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var loginTemplate = template.Must(template.New("login").Parse(`<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8">
				<title>Login</title>
			</head>

			<body>
				<form method="POST" action="/smartlights/login">
					<input type="hidden" name="state" value="{{ .State }}"/>
					Name:<br/>
					<input type="text" name="name"/><br/>
					Pass:<br/>
					<input type="password" name="pass"/><br/>
					<input type="submit"/>
				</form>
			</body>
		</html>`))

type loginData struct {
	State string
}

type userData struct {
	Name         string
	PasswordHash []byte // bcrypt
}

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalf("no API_KEY env set")
	}
	wantClientID := os.Getenv("CLIENT_ID")
	if wantClientID == "" {
		log.Fatalf("no CLIENT_ID env set")
	}
	wantRedirectURI := os.Getenv("REDIRECT_URI")
	if wantRedirectURI == "" {
		log.Fatalf("no REDIRECT_URI env set")
	}
	const wantResponseType = "token"
	parsedRedirectURI, err := url.Parse(wantRedirectURI)
	if err != nil {
		log.Fatalf("failed to parse REDIRECT_URI: %s", err)
	}
	rawUsers := os.Getenv("USERS")
	if rawUsers == "" {
		log.Fatalf("no USERS env set")
	}
	var users []userData
	if err := json.Unmarshal([]byte(rawUsers), &users); err != nil {
		log.Fatalf("failed to parse USERS env var: %s", err)
	}
	hashByName := make(map[string][]byte, len(users))
	for i := range users {
		user := &users[i]
		hashByName[user.Name] = user.PasswordHash
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/smartlights/login", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			query, err := url.ParseQuery(r.URL.RawQuery)
			if err != nil {
				log.Printf("login called with invalid query: %s", err)
				http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			if clientID := query.Get("client_id"); clientID != wantClientID {
				log.Printf("login called with invalid client_id %q, want %q", clientID, wantClientID)
				http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			redirectURI := query.Get("redirect_uri")
			if redirectURI != wantRedirectURI {
				log.Printf("got unauthorized redirect_uri %q, want %q", redirectURI, wantRedirectURI)
				http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			if responseType := query.Get("response_type"); responseType != wantResponseType {
				log.Printf("login called with invalid response_type %q, want %q", responseType, wantResponseType)
				http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			err = loginTemplate.Execute(rw, loginData{
				State: query.Get("state"),
			})
			if err != nil {
				log.Printf("error executing login template: %s", err)
				return
			}
		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				log.Printf("failed to login parse request formdata: %s", err)
				http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			var (
				state = r.PostForm.Get("state")
				name  = r.PostForm.Get("name")
				pass  = r.PostForm.Get("pass")
			)
			if name == "" || pass == "" {
				log.Printf("missing user/pass")
				http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			hash := hashByName[name]
			if hash == nil {
				log.Printf("unknown user %q", name)
				http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			if err := bcrypt.CompareHashAndPassword(hash, []byte(pass)); err != nil {
				log.Printf("wrong password for user %q: %s", name, err)
				http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			location := *parsedRedirectURI
			location.RawQuery = url.Values{
				"access_token": {name + ":" + apiKey},
				"token_type":   {"bearer"},
				"state":        {state},
			}.Encode()
			rw.Header().Set("Location", location.String())
			rw.WriteHeader(http.StatusFound)
		default:
			http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
	err = http.ListenAndServe("127.0.0.1:18917", mux)
	log.Println("finished serving with err =", err)
}
