package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const wantResponseType = "code" // Auth code flow used for Smart Home Actions

var loginTemplate = template.Must(template.New("login").Parse(`<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8">
				<title>Login</title>
			</head>

			<body>
				<form method="POST" action="{{ .LoginPath }}">
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
	LoginPath string
	State     string
}
type oauthServer struct {
	loginPath          string
	tokenPath          string
	key                *rsa.PrivateKey
	clientID           string
	clientSecret       string
	rawRedirectURL     string
	redirectURL        *url.URL
	userPasswordHashes map[string][]byte // bcrypt
}

func oauthServerFromEnv(loginPath, tokenPath string) (*oauthServer, error) {
	res := oauthServer{
		loginPath:      loginPath,
		tokenPath:      tokenPath,
		clientID:       os.Getenv("CLIENT_ID"),
		clientSecret:   os.Getenv("CLIENT_SECRET"),
		rawRedirectURL: os.Getenv("REDIRECT_URI"),
	}

	if res.clientID == "" {
		return nil, errors.New("no CLIENT_ID env set")
	}

	if res.clientSecret == "" {
		return nil, errors.New("no CLIENT_SECRET env set")
	}

	if res.rawRedirectURL == "" {
		return nil, errors.New("no REDIRECT_URI env set")
	}
	var err error
	res.redirectURL, err = url.Parse(res.rawRedirectURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse REDIRECT_URI: %s", err)
	}

	rawUsers := os.Getenv("USERS")
	if rawUsers == "" {
		return nil, errors.New("no USERS env set")
	}
	type userData struct {
		Name         string
		PasswordHash []byte // bcrypt
	}
	var users []userData
	if err := json.Unmarshal([]byte(rawUsers), &users); err != nil {
		return nil, fmt.Errorf("failed to parse USERS env var: %s", err)
	}
	res.userPasswordHashes = make(map[string][]byte, len(users))
	for i := range users {
		user := &users[i]
		res.userPasswordHashes[user.Name] = user.PasswordHash
	}

	pkPath := os.Getenv("PK_PATH")
	if pkPath == "" {
		return nil, errors.New("no PK_PATH env set")
	}
	pkFile, err := os.Open(pkPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open PK_PATH %q: %s", pkPath, err)
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if closeErr := pkFile.Close(); closeErr != nil { // always close regardless of err!
		return nil, fmt.Errorf("failed to close PK file %q: %s", pkPath, closeErr)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read PK file %q: %s", pkPath, err)
	}
	pkBlock, _ := pem.Decode(pkBytes)
	if pkBlock == nil {
		return nil, fmt.Errorf("PK file %q contains no PEM block", pkPath)
	}
	if pkBlock.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("unsupposed PK type %q", pkBlock.Type)
	}
	res.key, err = x509.ParsePKCS1PrivateKey(pkBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing RSA PK in %q: %s", pkPath, err)
	}

	return &res, nil
}

func (s *oauthServer) serveLogin(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Printf("login called with invalid query: %s", err)
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if clientID := query.Get("client_id"); clientID != s.clientID {
			log.Printf("login called with invalid client_id %q, want %q", clientID, s.clientID)
			http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		redirectURI := query.Get("redirect_uri")
		if redirectURI != s.rawRedirectURL {
			log.Printf("got unauthorized redirect_uri %q, want %q", redirectURI, s.rawRedirectURL)
			http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		if responseType := query.Get("response_type"); responseType != wantResponseType {
			log.Printf("login called with invalid response_type %q, want %q", responseType, wantResponseType)
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = loginTemplate.Execute(rw, loginData{
			LoginPath: s.loginPath,
			State:     query.Get("state"),
		})
		if err != nil {
			log.Printf("error executing login template: %s", err)
			return
		}

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Printf("failed to parse login request formdata: %s", err)
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
		hash := s.userPasswordHashes[name]
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
		authCode, err := s.signAuthToken(&authTokenPayload{
			Type:       typeAuthCode,
			User:       name,
			ClientID:   s.clientID,
			ValidUntil: time.Now().Add(10 * time.Minute),
		})
		if err != nil {
			log.Printf("failed to create auth code: %s", err)
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		location := *s.redirectURL
		location.RawQuery = url.Values{
			"code":  {authCode},
			"state": {state},
		}.Encode()
		rw.Header().Set("Location", location.String())
		rw.WriteHeader(http.StatusFound)
		log.Printf("successful login by %q for %q", name, s.clientID)

	default:
		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (s *oauthServer) serveToken(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Printf("failed to parse token request formdata: %s", err)
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		clientID := r.PostForm.Get("client_id")
		if clientID != s.clientID {
			log.Printf("token called with invalid client_id %q, want %q", clientID, s.clientID)
			http.Error(rw, `{"error": "invalid_grant"}`, http.StatusBadRequest)
			return
		}
		clientSecret := r.PostForm.Get("client_secret")
		if clientSecret != s.clientSecret {
			log.Printf("token called with invalid client_secret %q, want %q", clientSecret, s.clientSecret)
			http.Error(rw, `{"error": "invalid_grant"}`, http.StatusBadRequest)
			return
		}
		const secondsToLive = 60 * 60 * 24
		switch grantType := r.PostForm.Get("grant_type"); grantType {
		case "authorization_code":
			raw := r.PostForm.Get("code")
			if raw == "" {
				log.Printf("token authorization_code called without code")
				http.Error(rw, `{"error": "invalid_grant"}`, http.StatusBadRequest)
				return
			}
			token, err := s.parseAuthToken(raw, time.Now(), typeAuthCode, clientID)
			if err != nil {
				log.Printf("token authorization_code invalid: %s", err)
				http.Error(rw, `{"error": "invalid_grant"}`, http.StatusBadRequest)
				return
			}
			accessToken, err := s.signAuthToken(&authTokenPayload{
				Type:       typeAccessToken,
				User:       token.User,
				ClientID:   token.ClientID,
				ValidUntil: time.Now().Add(secondsToLive * time.Second),
			})
			if err != nil {
				log.Printf("token authorization_code failed to generate access token: %s", err)
				http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			refreshToken, err := s.signAuthToken(&authTokenPayload{
				Type:     typeRefreshToken,
				User:     token.User,
				ClientID: token.ClientID,
			})
			if err != nil {
				log.Printf("token authorization_code failed to generate refresh token: %s", err)
				http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			resp := struct {
				TokenType    string `json:"token_type"`
				AccessToken  string `json:"access_token"`
				RefreshToken string `json:"refresh_token"`
				ExpiresIn    int    `json:"expires_in"`
			}{
				TokenType:    "Bearer",
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				ExpiresIn:    secondsToLive,
			}
			if err := json.NewEncoder(rw).Encode(&resp); err != nil {
				log.Printf("token authorization_code writing response: %s", err)
				return
			}
			log.Printf("token authorization_code successful response for %q", clientID)
		case "refresh_token":
			raw := r.PostForm.Get("refresh_token")
			if raw == "" {
				log.Printf("token refresh_token called without refresh_token")
				http.Error(rw, `{"error": "invalid_grant"}`, http.StatusBadRequest)
				return
			}
			token, err := s.parseAuthToken(raw, time.Time{}, typeRefreshToken, clientID)
			if err != nil {
				log.Printf("token refresh_token invalid: %s", err)
				http.Error(rw, `{"error": "invalid_grant"}`, http.StatusBadRequest)
				return
			}
			accessToken, err := s.signAuthToken(&authTokenPayload{
				Type:       typeAccessToken,
				User:       token.User,
				ClientID:   token.ClientID,
				ValidUntil: time.Now().Add(secondsToLive * time.Second),
			})
			if err != nil {
				log.Printf("token refresh_token failed to generate access token: %s", err)
				http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			resp := struct {
				TokenType   string `json:"token_type"`
				AccessToken string `json:"access_token"`
				ExpiresIn   int    `json:"expires_in"`
			}{
				TokenType:   "Bearer",
				AccessToken: accessToken,
				ExpiresIn:   secondsToLive,
			}
			if err := json.NewEncoder(rw).Encode(&resp); err != nil {
				log.Printf("token refresh_token writing response: %s", err)
				return
			}
			log.Printf("token refresh_token successful response for %q", clientID)
		default:
			log.Printf("token called with invalid grant_type %q for %q", grantType, clientID)
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	default:
		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

type authTokenPayload struct {
	Type       authTokenType `json:"type"`
	User       string        `json:"user"`
	ClientID   string        `json:"client_id"`
	ValidUntil time.Time     `json:"valid_until"`
}

type authTokenType string

const (
	typeAuthCode     authTokenType = "auth_code"
	typeAccessToken  authTokenType = "access_token"
	typeRefreshToken authTokenType = "refresh_token"
)

func (s *oauthServer) signAuthToken(payload *authTokenPayload) (string, error) {
	j, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("json marshal: %s", err)
	}
	b64 := base64.StdEncoding.EncodeToString(j)
	sig, err := s.sign(j)
	if err != nil {
		return "", fmt.Errorf("signing: %s", err)
	}
	return b64 + "." + sig, nil
}

type authTokenParser interface {
	parseAuthToken(rawToken string, now time.Time, wantType authTokenType, wantClient string) (*authTokenPayload, error)
}

var _ authTokenParser = (*oauthServer)(nil)

func (s *oauthServer) parseAuthToken(rawToken string, now time.Time, wantType authTokenType, wantClient string) (*authTokenPayload, error) {
	pieces := strings.SplitN(rawToken, ".", 2)
	if len(pieces) != 2 {
		return nil, errors.New("token contains no '.'")
	}
	j, err := base64.StdEncoding.DecodeString(pieces[0])
	if err != nil {
		return nil, fmt.Errorf("decoding payload: %s", err)
	}
	if err := s.verify(j, pieces[1]); err != nil {
		return nil, fmt.Errorf("invalid signature: %s", err)
	}
	var res authTokenPayload
	if err = json.Unmarshal(j, &res); err != nil {
		return nil, fmt.Errorf("unmarshalling payload: %s", err)
	}
	if !now.IsZero() && res.ValidUntil.Before(now) {
		return nil, errors.New("token expired")
	}
	if res.Type != wantType {
		return nil, fmt.Errorf("token has unexpected type %q, want %q", res.Type, wantType)
	}
	if wantClient != "" && res.ClientID != wantClient {
		return nil, fmt.Errorf("caller %q doesn't match token holder %q", wantClient, res.ClientID)
	}
	return &res, nil
}

func (s *oauthServer) sign(data []byte) (string, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(data)); err != nil {
		return "", fmt.Errorf("hashing: %s", err)
	}
	d := h.Sum(nil)
	sig, err := rsa.SignPKCS1v15(rand.Reader, s.key, crypto.SHA256, d)
	if err != nil {
		return "", fmt.Errorf("signing: %s", err)
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}
func (s *oauthServer) verify(data []byte, signature64 string) error {
	sig, err := base64.StdEncoding.DecodeString(signature64)
	if err != nil {
		return fmt.Errorf("decoding: %s", err)
	}
	h := sha256.New()
	_, err = h.Write(data)
	if err != nil {
		return fmt.Errorf("hashing: %s", err)
	}
	d := h.Sum(nil)
	return rsa.VerifyPKCS1v15(&s.key.PublicKey, crypto.SHA256, d, sig)
}
