package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mrwonko/smartlights/config"
	"github.com/mrwonko/smartlights/internal/protocol"
)

// serveFulfillment returns an http.HandlerFunc that handles requests made by Google Home.
func serveFulfillment(tokenParser authTokenParser, pc *pubsubClient, sc *stateCache) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		token, status, err := checkFulfillmentAuth(tokenParser, r)
		if err != nil {
			log.Print(err)
			http.Error(rw, http.StatusText(status), status)
			return
		}
		req, status, err := parseFulfillmentRequest(r)
		if err != nil {
			log.Print(err)
			http.Error(rw, http.StatusText(status), status)
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
			serveFulfillmentSync(req, token, rw)
		case intentDisconnect:
			serveFulfillmentDisconnect(rw)
		case intentExecute:
			serveFulfillmentExecute(r.Context(), pc, req, input, rw)
		case intentQuery:
			serveFulfillmentQuery(r.Context(), sc, req, input, pc.Report, rw)
		default:
			log.Printf("fulfillment unsupported intent %q", input.Intent)
			http.Error(rw, "unsupported intent", http.StatusNotImplemented)
		}
	}
}

func checkFulfillmentAuth(tokenParser authTokenParser, r *http.Request) (*authTokenPayload, int, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil, http.StatusUnauthorized, errors.New("fulfillment called without auth")
	}
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(auth, bearerPrefix) {
		return nil, http.StatusUnauthorized, errors.New("fulfillment called with non-bearer auth")
	}
	auth = strings.TrimPrefix(auth, bearerPrefix)
	token, err := tokenParser.parseAuthToken(auth, time.Now(), typeAccessToken, "")
	if err != nil {
		return nil, http.StatusForbidden, fmt.Errorf("fulfillment called with invalid token: %s", err)
	}
	return token, http.StatusOK, nil
}

func parseFulfillmentRequest(r *http.Request) (*request, int, error) {
	if r.Method != http.MethodPost {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("fulfillment called with invalid method %q", r.Method)
	}
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("fulfillment failed to parse body: %s", err)
	}
	return &req, http.StatusOK, nil
}

func serveFulfillmentSync(req *request, token *authTokenPayload, rw http.ResponseWriter) {
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
}

func serveFulfillmentDisconnect(rw http.ResponseWriter) {
	log.Printf("TODO implement disconnect")
	http.Error(rw, "TODO implement disconnect", http.StatusNotImplemented)
}

func serveFulfillmentExecute(ctx context.Context, pc *pubsubClient, req *request, input *requestInput, rw http.ResponseWriter) {
	inputPayload := requestPayloadExecute{}
	if err := json.Unmarshal(input.Payload, &inputPayload); err != nil {
		log.Printf("fulfillment execute json parse: %s", err)
		http.Error(rw, "failed to parse body", http.StatusBadRequest)
		return
	}

	messagesByPi, errCode, err := convertExecuteCommand(&inputPayload)
	if err != nil {
		log.Printf("fulfilment execute convert: %s", err)
		if err := json.NewEncoder(rw).Encode(response{
			RequestID: req.RequestID,
			Payload: responsePayloadError{
				ErrorCode: errCode,
			},
		}); err != nil {
			log.Printf("fulfillment execute error response: %s", err)
		}
		return
	}

	type result struct {
		pi  int
		err error
	}
	results := make([]result, len(messagesByPi))
	i := 0
	var wg sync.WaitGroup
	for pi, cmd := range messagesByPi {
		wg.Add(1)
		go func(i int, pi int, cmd protocol.ExecuteMessage) {
			defer wg.Done()
			err := pc.Execute(ctx, pi, cmd)
			results[i] = result{pi: pi, err: err}
		}(i, pi, cmd)
		i++
	}
	wg.Wait()

	// for every requested device, contain the list of errors, which may be empty
	deviceErrors := map[config.ID][]error{}
	for _, res := range results {
		ids := map[config.ID]struct{}{}
		for _, cmd := range messagesByPi[res.pi].Commands {
			for _, id := range cmd.Devices {
				ids[id] = struct{}{}
			}
		}
		var errs []error
		if res.err != nil {
			errs = []error{res.err}
		}
		for id := range ids {
			deviceErrors[id] = append(deviceErrors[id], errs...)
		}
	}

	payload := responsePayloadExecute{
		Commands: make([]responsePayloadExecuteCommand, len(deviceErrors)),
	}
	curCommand := 0
	for id, errs := range deviceErrors {
		cmd := &payload.Commands[curCommand]
		curCommand++
		cmd.IDs = []string{strconv.Itoa(int(id))}
		if len(errs) == 0 {
			cmd.Status = statusPending
		} else {
			var b strings.Builder
			for i, err := range errs {
				if i > 0 {
					b.WriteString(", ")
				}
				b.WriteString(err.Error())
			}
			msg := b.String()
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
	log.Printf("successful execute")
}

func convertExecuteCommand(ex *requestPayloadExecute) (map[int]protocol.ExecuteMessage, errorCode, error) {
	res := map[int]protocol.ExecuteMessage{}
	for _, cmd := range ex.Commands {
		devicesByPi := map[int][]config.ID{}
		for _, d := range cmd.Devices {
			id, err := strconv.Atoi(d.ID)
			if err != nil {
				return nil, errorCodeDeviceNotFound, fmt.Errorf("device ID %q parse: %s", d.ID, err)
			}
			l := config.Lights[config.ID(id)]
			if l == nil {
				return nil, errorCodeDeviceNotFound, fmt.Errorf("unknown device ID %d", id)
			}
			devicesByPi[l.Pi] = append(devicesByPi[l.Pi], config.ID(id))
		}
		executions := make([]protocol.ExecuteExecution, len(cmd.Execution))
		for i, ex := range cmd.Execution {
			switch ex.Command {
			case "action.devices.commands.OnOff":
				on, ok := ex.Params["on"].(bool)
				if !ok {
					return nil, errorCodeProtocolError, fmt.Errorf(`OnOff command without "on" param`)
				}
				executions[i] = protocol.ExecuteExecutionOnOff{On: on}
			default:
				return nil, errorCodeFunctionNotSupported, fmt.Errorf("unsupported command %q", ex.Command)
			}
		}
		for pi, devices := range devicesByPi {
			res[pi] = protocol.ExecuteMessage{Commands: append(res[pi].Commands, &protocol.ExecuteCommand{
				Devices:    devices,
				Executions: executions,
			})}
		}
	}
	return res, "", nil
}

func serveFulfillmentQuery(ctx context.Context, sc *stateCache, req *request, input *requestInput, requestReport func(ctx context.Context, pi int, reason string) error, rw http.ResponseWriter) {
	// QUERY intent https://developers.google.com/assistant/smarthome/reference/intent/query
	// requests state of specific devices
	inputPayload := requestPayloadQuery{}
	if err := json.Unmarshal(input.Payload, &inputPayload); err != nil {
		log.Printf("fulfillment query json parse: %s", err)
		http.Error(rw, "failed to parse body", http.StatusBadRequest)
		return
	}
	log.Printf("fulfiment query request: %v", inputPayload)
	devices := map[string]map[string]interface{}{}
	deviceStates := sc.get()
	reportsRequestedFromPis := map[int]struct{}{}
	for _, cur := range inputPayload.Devices {
		id := cur.ID
		if _, ok := devices[id]; ok {
			// duplicate, already processed
			continue
		}
		states, ok := deviceStates[id]
		if !ok {
			// no state cached for this device yet
			// request a report, but for now, report as offline
			parsedID, err := strconv.Atoi(id)
			if err != nil {
				log.Printf("received query for non-int device %q", id)
			} else if light, ok := config.Lights[config.ID(parsedID)]; !ok {
				log.Printf("received query for unknown device %d", parsedID)
			} else if _, requested := reportsRequestedFromPis[light.Pi]; !requested {
				if err := requestReport(ctx, light.Pi, "query"); err != nil {
					log.Printf("failed to request report from pi %d: %s", light.Pi, err)
				}
				reportsRequestedFromPis[light.Pi] = struct{}{}
			}
			devices[id] = map[string]interface{}{
				"online": false,
				"status": "OFFLINE",
			}
			continue
		}
		dev := map[string]interface{}{
			"online": true,
			"status": "SUCCESS",
		}
		if states.OnOff != nil {
			dev[string(stateOn)] = states.OnOff.On
		}
		devices[id] = dev
	}
	payload := responsePayloadQuery{
		Devices: devices,
	}
	if err := json.NewEncoder(io.MultiWriter(rw, os.Stderr)).Encode(response{
		RequestID: req.RequestID,
		Payload:   payload,
	}); err != nil {
		log.Printf("fulfillment query send response: %s", err)
		return
	}
	log.Printf("successful query")
}
