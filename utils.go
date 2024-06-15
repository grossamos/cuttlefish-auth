package main

const GOOGLE_STUN_SERVER = "stun:stun.l.google.com:19302"
const DEVICE_PARTNER string = "device_partner"
const CLIENT_PARTNER string = "client_partner"

func CreateGoogleStunSignaling() ConfigMessage {
  configMsg := NewConfigMessage([]map[string]string{
		{"url": GOOGLE_STUN_SERVER},
	})
	return configMsg
}
