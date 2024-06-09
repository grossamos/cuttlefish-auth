package model

type RegisterMessage struct {
	MessageType string      `json:"message_type"`
	DeviceID    string      `json:"device_id"`
	DeviceInfo  interface{} `json:"device_info"`
}

type ForwardMessage struct {
	MessageType string      `json:"message_type"`
	ClientID    int         `json:"client_id"`
	Payload     interface{} `json:"payload"`
}
