package config

type ExecuteMessage struct {
	GPIO uint8
	On   bool
}

type QueryRequestMessage struct {
	RequestID string
	IDs       []ID
}

type QueryResponseMessage struct {
	RequestID string
	// TODO
}
