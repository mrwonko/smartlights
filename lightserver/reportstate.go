package main

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// reportState implements https://developers.google.com/assistant/smarthome/develop/report-state
func reportState(ctx context.Context, client *http.Client, agentUserID string, serviceAccountKey *rsa.PrivateKey, serviceAccountEmail string) error { // TODO: state

	// create JWT to request access token with
	// TODO: cache response until it expires, recreate on demand
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":   serviceAccountEmail,
		"scope": "https://www.googleapis.com/auth/homegraph",
		"aud":   "https://accounts.google.com/o/oauth2/token",
		"iat":   now.Unix(),
		"exp":   now.Add(time.Hour).Unix(),
	})
	jwt, err := token.SignedString(serviceAccountKey)
	if err != nil {
		return fmt.Errorf("signing JWT: %w", err)
	}
	// exchange for access token
	data := url.Values{}
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	data.Set("assertion", jwt)
	req, err := http.NewRequest(http.MethodPost, "https://accounts.google.com/o/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("creating JWT to Access Token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("doing JWT to Access Token request: %w", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if closeErr := resp.Body.Close(); closeErr != nil {
		log.Printf("closing JWT to Access Token response body: %s", err)
		// non-fatal
	}
	if resp.StatusCode != http.StatusOK {
		// see https://developers.google.com/actions/smarthome/report-state#error_responses
		return fmt.Errorf("response status %d for JWT to Access Token request %q: %q", resp.StatusCode, data.Encode(), respBody)
	}
	var parsedResp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"` // seconds
	}
	if err = json.Unmarshal(respBody, &parsedResp); err != nil {
		return fmt.Errorf("parsing JWT to Access Token response body: %w", err)
	}

	reqBody, err := json.Marshal(struct {
		RequestID   string              `json:"requestId"`
		AgentUserID string              `json:"agentUserId"`
		Payload     requestPayloadState `json:"payload"`
	}{
		RequestID:   uuid.New().String(), // TODO
		AgentUserID: agentUserID,
		Payload:     requestPayloadState{
			// TODO
		},
	})
	if err != nil {
		return fmt.Errorf("request marshalling failed: %w", err)
	}
	req, err = http.NewRequest(http.MethodPost,
		(&url.URL{
			Scheme: "https",
			Host:   "homegraph.googleapis.com",
			Path:   "v1/devices:reportStateAndNotification",
		}).String(),
		bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+parsedResp.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	respBody, err = ioutil.ReadAll(resp.Body)
	if closeErr := resp.Body.Close(); closeErr != nil {
		log.Printf("state response body closing failed: %s", err)
		// non-fatal
	}
	if err != nil {
		return fmt.Errorf("response body read failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		// see https://developers.google.com/actions/smarthome/report-state#error_responses
		return fmt.Errorf("response status %d for request `%s`: %q", resp.StatusCode, reqBody, respBody)
	}
	return nil
}