package main

import (
	"encoding/json"
	"errors"
	"log"
)

const REGISTER_MESSAGE_TYPE = "register"
const CONNECT_MESSAGE_TYPE = "register"
const FORWARD_MESSAGE_TYPE = "forward"

func HandleSignaling(partner string, message string) (string, error) {
  var msgTypeContainer mesageTypeContainer

  if msgTypeContainer.MessageType == REGISTER_MESSAGE_TYPE || msgTypeContainer.MessageType == CONNECT_MESSAGE_TYPE {
    deviceType, err := getDeviceType(message)

    response, err := json.Marshal(deviceType)
    if err != nil {
      log.Println("Failed to marshal cfg response", err)
      return "", err
    }
    return string(response), nil

  } else if msgTypeContainer.MessageType == FORWARD_MESSAGE_TYPE {
    return "", nil
  } else {
    log.Println("Unkown message_type: ", msgTypeContainer.MessageType)
    return "", errors.New("unknown message type: " + msgTypeContainer.MessageType)
  }
}


func getDeviceType(message string) (interface{}, error) {
    var deviceInfoContainer deviceInfoContainerType

    err := json.Unmarshal([]byte(message), &deviceInfoContainer)
    if err != nil {
      log.Println("Unable to unmarshal forward json:", err)
      return nil, err
    }

  return deviceInfoContainer.DeviceInfo, err

}

type deviceInfoContainerType struct {
  DeviceInfo interface{} `json:"device_info"`
}

