package protocol

type (
	StateMessage struct {
		Devices map[string]DeviceStates
	}
	DeviceStates struct {
		OnOff *OnOffState `json:",omitempty"`
	}
	OnOffState struct {
		On bool
	}
)
