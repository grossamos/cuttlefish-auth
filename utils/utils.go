package utils

import (
	"errors"

	"github.com/grossamos/cuttlefish-auth/models"
)

const GOOGLE_STUN_SERVER = "stun:stun.l.google.com:19302"
const DEVICE_PARTNER string = "device_partner"
const CLIENT_PARTNER string = "client_partner"

func CreateGoogleStunSignalingConfigMessage() models.ConfigMessage {
  urls := make([]string, 1)
  urls[0] = GOOGLE_STUN_SERVER
  configMsg := models.NewConfigMessage([]map[string][]string{
		{"urls": urls},
	})
	return configMsg
}

func GetIdOfChannel(databank *map[string]models.DataBankEntry, ch chan models.ChannelMessage) (uint, error) {
  for _, entry := range *databank {
    for id, client := range entry.Clients {
      if client == ch {
        return id, nil
      }
    }
  }
  return 0, errors.New("Client Channel doesn't exist")
}

func GetDeviceOfChannel(databank *map[string]models.DataBankEntry, ch chan models.ChannelMessage) (string, error) {
  for deviceId, entry := range *databank {
    if entry.Channel == ch {
      return deviceId, nil
    }
  }
  return "", errors.New("Device Channel doesn't exist")
}

func GetDeviceOfClient(databank *map[string]models.DataBankEntry, clientID uint) (string, error) {
  for deviceId, entry := range *databank {
    for id := range entry.Clients {
      if id == clientID {
        return deviceId, nil
      }
    }
  }
  return "", errors.New("Client ID doesn't exist")
}
