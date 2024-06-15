package main

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

type ConfigMessage struct {
	MessageType string            `json:"message_type"`
  IceServers  []map[string]string `json:"ice_servers"`
}

func NewConfigMessage(iceServers []map[string]string) ConfigMessage {
  return ConfigMessage{
    MessageType: "config",
    IceServers: iceServers,
  }
}

