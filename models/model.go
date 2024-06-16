package models

type RegisterMessage struct {
	MessageType string      `json:"message_type"`
	DeviceID    string      `json:"device_id"`
	DeviceInfo  interface{} `json:"device_info"`
}

type ForwardMessage struct {
	MessageType string      `json:"message_type"`
	Payload     interface{} `json:"payload"`
}

type ClientMessage struct {
	MessageType string `json:"message_type"`
	ClientID    uint   `json:"client_id"`
	Payload     interface{} `json:"payload"`
}

type DeviceMessage struct {
	MessageType string `json:"message_type"`
	DeviceID    string      `json:"device_id"`
	Payload     interface{} `json:"payload"`
}

type DeviceInfoMessage struct {
	MessageType string      `json:"message_type"`
	DeviceInfo  interface{} `json:"device_info"`
}

type ConfigMessage struct {
	MessageType string              `json:"message_type"`
	IceServers  []map[string][]string `json:"ice_servers"`
}

func NewConfigMessage(iceServers []map[string][]string) ConfigMessage {
	return ConfigMessage{
		MessageType: "config",
		IceServers:  iceServers,
	}
}

type DataBankEntry struct {
	Channel    chan ChannelMessage
	DeviceInfo interface{}
	Clients    map[uint]chan ChannelMessage
}

type Client struct {
  Ch chan ChannelMessage
  ClientId uint
}

type ChannelMessage struct {
	ForwardYesWsNo bool
	Message        string
}
