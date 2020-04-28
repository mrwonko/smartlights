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
)

// serveFulfillment returns an http.HandlerFunc that handles requests made by Google Home.
func serveFulfillment(tokenParser authTokenParser, pc *pubsubClient) http.HandlerFunc {
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
			serveFulfillmentQuery(rw)
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

type deviceErr struct {
	device config.ID
	err    error
}

func serveFulfillmentExecute(ctx context.Context, pc *pubsubClient, req *request, input *requestInput, rw http.ResponseWriter) {
	inputPayload := requestPayloadExecute{}
	if err := json.Unmarshal(input.Payload, &inputPayload); err != nil {
		log.Printf("fulfillment execute json parse: %s", err)
		http.Error(rw, "failed to parse body", http.StatusBadRequest)
		return
	}
	// write individual errors to deviceErrs. Once it is closed, a map of all collected errors is sent to deviceErrsMapChan.
	deviceErrs, deviceErrsMapChan := func(ctx context.Context) (chan<- deviceErr, <-chan map[config.ID][]error) {
		reqChan := make(chan deviceErr)
		resChan := make(chan map[config.ID][]error)
		go func(ctx context.Context, reqChan <-chan deviceErr, resChan chan<- map[config.ID][]error) {
			defer close(resChan)
			res := map[config.ID][]error{}
			for {
				select {
				case req, ok := <-reqChan:
					if !ok {
						resChan <- res
						return
					}
					errs := res[req.device]
					if req.err != nil {
						errs = append(errs, req.err)
					}
					res[req.device] = errs
				case <-ctx.Done():
					return
				}
			}
		}(ctx, reqChan, resChan)
		return reqChan, resChan
	}(ctx)
	var wg sync.WaitGroup
	for i := range inputPayload.Commands {
		command := &inputPayload.Commands[i]
		if errCode, err := handleCommand(ctx, pc, &wg, deviceErrs, command); err != nil {
			log.Printf("fulfillment execute error response: %s", err)
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
	}
	wg.Wait()
	close(deviceErrs)
	deviceErrsMap := <-deviceErrsMapChan

	type failedDevice struct {
		id   config.ID
		errs []error
	}
	var successfulDevices []config.ID
	var failedDevices []failedDevice
	for id, errs := range deviceErrsMap {
		if len(errs) == 0 {
			successfulDevices = append(successfulDevices, id)
		} else {
			failedDevices = append(failedDevices, failedDevice{id: id, errs: errs})
		}
	}

	numCommands := 0
	if len(successfulDevices) > 0 {
		numCommands++
	}
	if len(failedDevices) > 0 {
		numCommands++
	}

	payload := responsePayloadExecute{
		Commands: make([]responsePayloadExecuteCommand, numCommands),
	}
	curCommand := 0
	if len(successfulDevices) > 0 {
		cmd := &payload.Commands[curCommand]
		curCommand++
		cmd.IDs = make([]string, len(successfulDevices))
		for i, id := range successfulDevices {
			cmd.IDs[i] = strconv.Itoa(int(id))
		}
		cmd.Status = statusPending
	}
	if len(failedDevices) > 0 {
		cmd := &payload.Commands[curCommand]
		curCommand++
		cmd.IDs = make([]string, len(failedDevices))
		var b strings.Builder
		for i, fd := range failedDevices {
			if i > 0 {
				b.WriteString(", ")
			}
			cmd.IDs[i] = strconv.Itoa(int(fd.id))
			fmt.Fprintf(&b, "error(s) toggling device %d:", fd.id)
			for _, err := range fd.errs {
				fmt.Fprintf(&b, " %q", err)
			}
		}
		msg := b.String()
		cmd.DebugString = &msg
		cmd.Status = statusError
		cmd.ErrorCode = errorCodeTransientError
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

// handleCommand sends the results of handling each command to the deviceErrs channel asynchronously, incrementing wg until that is done.
func handleCommand(ctx context.Context, pc *pubsubClient, wg *sync.WaitGroup, deviceErrs chan<- deviceErr, command *requestPayloadExecuteCommand) (errorCode, error) {
	devices := make(map[config.ID]*config.Light, len(command.Devices))
	for i := range command.Devices {
		id, err := strconv.Atoi(command.Devices[i].ID)
		if err != nil {
			return errorCodeDeviceNotFound, fmt.Errorf("fulfillment execute device ID %q parse: %s", command.Devices[i].ID, err)
		}
		light := config.Lights[config.ID(id)]
		if light == nil {
			return errorCodeDeviceNotFound, fmt.Errorf("unknown device %d", id)
		}
		devices[config.ID(id)] = light
	}

	ons := make([]bool, 0, len(command.Execution)) // TODO: remove
	for i := range command.Execution {
		e := &command.Execution[i]
		switch e.Command {
		case "action.devices.commands.OnOff":
			on, ok := e.Params["on"].(bool)
			if !ok {
				return errorCodeProtocolError, fmt.Errorf(`fulfillment execute OnOff command without "on" param`)
			}
			ons = append(ons, on)
		default:
			return errorCodeFunctionNotSupported, fmt.Errorf("unsupported command %q", e.Command)
		}
	}
	// FIXME: collect errors properly so this limit can be removed
	if len(ons) > 1 {
		return errorCodeProtocolError, fmt.Errorf("too many ons: %v", ons)
	}

	for _, on := range ons {
		for id := range devices {
			wg.Add(1)
			go func(id config.ID, on bool) {
				deviceErrs <- deviceErr{
					device: id,
					err:    pc.OnOff(ctx, id, on),
				}
				wg.Done()
			}(id, on)
		}
	}
	return "", nil
}

func serveFulfillmentQuery(rw http.ResponseWriter) {
	log.Printf("TODO implement query")
	http.Error(rw, "TODO implement query", http.StatusNotImplemented)
}
