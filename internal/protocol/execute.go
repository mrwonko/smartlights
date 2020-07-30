package protocol

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mrwonko/smartlights/config"
)

type ExecuteMessage struct {
	Commands []*ExecuteCommand
}

type ExecuteCommand struct {
	Devices    []config.ID
	Executions []ExecuteExecution
}

type ExecuteExecution interface {
	execution() // this is a sum type, all possible implementations are in this package
}

var executeExecutions = map[string]reflect.Type{
	"onoff": reflect.TypeOf(ExecuteExecutionOnOff{}),
}

type ExecuteExecutionOnOff struct {
	On bool
}

func (ExecuteExecutionOnOff) execution() {}

func (ec *ExecuteCommand) MarshalJSON() ([]byte, error) {
	type (
		executionWireFormat struct {
			Type  string
			Value ExecuteExecution
		}
		wireFormat struct {
			Devices    []config.ID
			Executions []executionWireFormat
		}
	)
	convertedExecutions := make([]executionWireFormat, len(ec.Executions))
	for i, ee := range ec.Executions {
		convertedExecutions[i].Value = ee
		eeType := reflect.TypeOf(ee)
		for name, t := range executeExecutions {
			if t == eeType {
				convertedExecutions[i].Type = name
				break
			}
		}
		if convertedExecutions[i].Type == "" {
			return nil, fmt.Errorf("unhandled Execution type %s", eeType)
		}
	}
	return json.Marshal(wireFormat{
		Devices:    ec.Devices,
		Executions: convertedExecutions,
	})
}

func (ec *ExecuteCommand) UnmarshalJSON(data []byte) error {
	type (
		executionWireFormat struct {
			Type  string
			Value json.RawMessage
		}
		wireFormat struct {
			Devices    []config.ID
			Executions []executionWireFormat
		}
	)
	var wf wireFormat
	if err := json.Unmarshal(data, &wf); err != nil {
		return err
	}
	convertedExecutions := make([]ExecuteExecution, len(wf.Executions))
	for i, ee := range wf.Executions {
		t, ok := executeExecutions[ee.Type]
		if !ok {
			return fmt.Errorf("unhandled Execution type %q", ee.Type)
		}
		v := reflect.New(t)
		if err := json.Unmarshal([]byte(ee.Value), v.Interface()); err != nil {
			return err
		}
		convertedExecutions[i] = v.Elem().Interface().(ExecuteExecution)
	}
	ec.Devices = wf.Devices
	ec.Executions = convertedExecutions
	return nil
}
