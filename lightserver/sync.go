package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func sendSyncRequests(ctx context.Context, client *http.Client, trigger <-chan struct{}, googleAPIKey, agentUserID string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case _, ok := <-trigger:
			if !ok {
				return nil
			}
			reqBody, err := json.Marshal(struct {
				AgentUserID string `json:"agentUserId"`
			}{
				AgentUserID: agentUserID,
			})
			if err != nil {
				return fmt.Errorf("sync request marshalling failed: %w", err)
			}
			req, err := http.NewRequest(http.MethodPost,
				(&url.URL{
					Scheme: "https",
					Host:   "homegraph.googleapis.com",
					Path:   "v1/devices:requestSync",
					RawQuery: url.Values{
						"key": {googleAPIKey},
					}.Encode(),
				}).String(),
				bytes.NewReader(reqBody))
			if err != nil {
				return fmt.Errorf("sync request creation failed: %w", err)
			}
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("sync request failed: %s", err)
				continue
			}
			respBody, err := ioutil.ReadAll(resp.Body)
			if closeErr := resp.Body.Close(); closeErr != nil {
				log.Printf("sync request body closing failed: %s", err)
				// non-fatal
			}
			if err != nil {
				log.Printf("sync request body read failed: %s", err)
				continue
			}
			if resp.StatusCode != http.StatusOK {
				// see https://developers.google.com/actions/smarthome/report-state#error_responses
				log.Printf("sync request response status %d for request `%s`: %q", resp.StatusCode, reqBody, respBody)
				continue
			}
			log.Printf("sync request successful: %q", respBody)
		}
	}
}
