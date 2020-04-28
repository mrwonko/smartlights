package main

import (
	"encoding/json"
	"strconv"

	"github.com/mrwonko/smartlights/config"
)

type (
	request struct {
		RequestID string         `json:"requestId"`
		Inputs    []requestInput `json:"inputs"`
	}

	requestInput struct {
		Intent  requestIntent   `json:"intent"`
		Payload json.RawMessage `json:"payload"` // type depends on intent
	}

	requestIntent string

	requestPayloadQuery struct {
		Devices []struct {
			ID string `json:"id"`
		} `json:"devices"`
	}

	requestPayloadExecute struct {
		Commands []requestPayloadExecuteCommand `json:"commands"`
	}

	requestPayloadExecuteCommand struct {
		Devices []struct {
			ID string `json:"id"`
		} `json:"devices"`
		Execution []struct {
			Command string                 `json:"command"`
			Params  map[string]interface{} `json:"params"` // string/number/boolean
		} `json:"execution"`
	}

	response struct {
		RequestID string      `json:"requestId"`
		Payload   interface{} `json:"payload"`
	}

	responsePayloadError struct {
		ErrorCode   errorCode `json:"errorCode"`
		DebugString *string   `json:"debugString,omitempty"`
	}

	responsePayloadSync struct {
		AgentUserID string                      `json:"agentUserId"`
		Devices     []responsePayloadSyncDevice `json:"devices"`
	}

	responsePayloadSyncDevice struct {
		ID              string                        `json:"id"`
		Type            deviceType                    `json:"type"`
		Traits          []deviceTrait                 `json:"traits"`
		Name            responsePayloadSyncDeviceName `json:"name"`
		WillReportState bool                          `json:"willReportState"` // true = push, false = poll
		// omitted: roomHint
		// omitted: deviceInfo
		Attributes map[string]interface{} `json:"attributes,omitempty"`
	}

	responsePayloadSyncDeviceName struct {
		DefaultNames []string `json:"defaultNames,omitempty"` // manufacturer names
		Name         string   `json:"name"`                   // user-provided main name
		Nicknames    []string `json:"nicknames,omitempty"`    // user-provided nicknames
	}

	responsePayloadQuery struct {
		Devices map[string]map[string]interface{} `json:"devices"` // TODO see https://developers.google.com/actions/smarthome/report-state
	}

	responsePayloadExecute struct {
		Commands []responsePayloadExecuteCommand `json:"commands"`
	}

	responsePayloadExecuteCommand struct {
		IDs    []string      `json:"ids"`
		Status executeStatus `json:"status"`
		responsePayloadError
		States map[string]interface{} `json:"states,omitempty"`
	}

	requestPayloadState struct {
		// TODO
	}
)

type deviceType string

const (
	intentSync       requestIntent = "action.devices.SYNC"
	intentQuery      requestIntent = "action.devices.QUERY"
	intentExecute    requestIntent = "action.devices.EXECUTE"
	intentDisconnect requestIntent = "action.devices.DISCONNECT"
)

type deviceTrait string

const (
	typeLight deviceType = "action.devices.types.LIGHT"
)

type executeStatus string

const (
	statusSuccess executeStatus = "SUCCESS"
	statusPending executeStatus = "PENDING"
	statusOffline executeStatus = "OFFLINE"
	statusError   executeStatus = "ERROR"
)

type errorCode string

const (
	errorCodeAuthFailure          errorCode = "authFailure"
	errorCodeAuthExpired          errorCode = "authExpired"
	errorCodeDeviceOffline        errorCode = "deviceOffline"
	errorCodeTimeout              errorCode = "timeout"
	errorCodeTransientError       errorCode = "transientError"
	errorCodeDeviceNotFound       errorCode = "deviceNotFound"
	errorCodeProtocolError        errorCode = "protocolError"        // request processing failed
	errorCodeFunctionNotSupported errorCode = "functionNotSupported" // not implemented
	// full list at https://developers.google.com/actions/reference/smarthome/errors-exceptions
)

const (
	traitBrightness   deviceTrait = "action.devices.traits.Brightness"
	traitColorSetting deviceTrait = "action.devices.traits.ColorSetting"
	traitOnOff        deviceTrait = "action.devices.traits.OnOff"
	traitToggles      deviceTrait = "action.devices.traits.Toggles"
)

var devices = func() []responsePayloadSyncDevice {
	res := make([]responsePayloadSyncDevice, len(config.Lights))
	i := 0
	for id, light := range config.Lights {
		d := &res[i]
		i++
		d.ID = strconv.Itoa(int(id))
		d.Name.Name = light.Name
		d.Type = typeLight
		d.Traits = []deviceTrait{traitOnOff}
		d.WillReportState = false // TODO report state
		d.Attributes = map[string]interface{}{
			"commandOnlyOnOff": true, // TODO report state
		}
	}
	return res
}()
