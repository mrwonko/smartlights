package protocol

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mrwonko/smartlights/config"
)

func TestExecuteMessage_MarshalJSON(t *testing.T) {
	tests := []struct {
		name  string
		value ExecuteMessage
	}{
		{
			name: "christmas tree",
			value: ExecuteMessage{
				Commands: []*ExecuteCommand{
					{
						Devices: []config.ID{1, 2, 3},
						Executions: []ExecuteExecution{
							ExecuteExecutionOnOff{On: true},
							ExecuteExecutionOnOff{On: false},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := json.Marshal(tt.value)
			if err != nil {
				t.Fatal(err)
			}
			var decoded ExecuteMessage
			if err = json.Unmarshal(encoded, &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.value, decoded) {
				t.Errorf("got %#v, want %#v", decoded, tt.value)
			}
		})
	}
}
