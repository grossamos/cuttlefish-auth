package endpoints

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/grossamos/cuttlefish-auth/models"
	"github.com/grossamos/cuttlefish-auth/utils"
)

const REGISTER_MESSAGE_TYPE = "register"
const CONNECT_MESSAGE_TYPE = "connect"
const FORWARD_MESSAGE_TYPE = "forward"

func HandleRegistration(message string, dataBank *map[string]models.DataBankEntry, communicationPartner string, ch chan models.ChannelMessage) (string, error) {
  var response string
   
  if communicationPartner == utils.DEVICE_PARTNER {
    deviceInfo, deviceId, err := getDeviceDetails(message) // todo actually save the device type in the databank
    if err != nil {
      return "", err
    }
    dataBankEntry := models.DataBankEntry{Channel: ch, DeviceInfo: deviceInfo, Clients: make(map[uint]chan models.ChannelMessage, 0) }
    (*dataBank)[deviceId] = dataBankEntry
  } else {
    // TODO: also support client first connection
    deviceID, err := getClientDetails(message)
    if err != nil {
      return "", err
    }
    deviceDataBankEntry, ok := (*dataBank)[deviceID]
    if !ok {
      log.Println("Client connected before device")
      return "", errors.New("Client connected before device")
    }
    id := uint(len(deviceDataBankEntry.Clients))
    deviceDataBankEntry.Clients[id] = ch
    (*dataBank)[deviceID] = deviceDataBankEntry
  }
  cfgMsg := utils.CreateGoogleStunSignalingConfigMessage()
  responseRaw, err := json.Marshal(cfgMsg)
  if err != nil {
    log.Println("Failed to marshal cfg response", err)
    return "", err
  }

  response = string(responseRaw)
  log.Println(response)

  return string(response), nil
}


func getDeviceDetails(message string) (interface{}, string, error) {
  var deviceInfoContainer deviceDetailsContainerType

  err := json.Unmarshal([]byte(message), &deviceInfoContainer)
  if err != nil {
    log.Println("Unable to unmarshal register json:", err)
    return nil, "", err
  }

  return deviceInfoContainer.DeviceInfo, deviceInfoContainer.DeviceID, err

}

type deviceDetailsContainerType struct {
  DeviceInfo interface{} `json:"device_info"`
  DeviceID string `json:"device_id"`
}

func getClientDetails(message string) (string, error) {
  var clientDetailsContainer clientDetailsContainerType

  err := json.Unmarshal([]byte(message), &clientDetailsContainer)
  if err != nil {
    log.Println("Unable to unmarshal connect json:", err)
    return "", err
  }

  return clientDetailsContainer.DeviceID, nil
}

type clientDetailsContainerType struct {
  DeviceID string `json:"device_id"`
}

func HandleDeviceInfoPropagation(message string, dataBank *map[string]models.DataBankEntry) (string, error) {
  var clientDetailsContainer clientDetailsContainerType

  err := json.Unmarshal([]byte(message), &clientDetailsContainer)
  if err != nil {
    log.Println("Unable to unmarshal connect json:", err)
    return "", err
  }

  deviceID := clientDetailsContainer.DeviceID
  deviceInfo := (*dataBank)[deviceID].DeviceInfo

  response, err := json.Marshal(models.DeviceInfoMessage{MessageType: "device_info", DeviceInfo: deviceInfo})
  if err != nil {
    log.Println("Failed to marshal device info:", err)
    return "", err
  }

  return string(response), nil
}


func ForwardMessageFromSender(dataBank *map[string]models.DataBankEntry, message string, connectionPartner string, ch chan models.ChannelMessage) (error) {
  log.Println("TIME TO FORWARD")
  var fwdMsg models.ForwardMessage
  var responseObj interface{}
  var recipient chan models.ChannelMessage

  err := json.Unmarshal([]byte(message), &fwdMsg)
  if err != nil {
    log.Println("Unable to unmarshal connect json:", err)
    return err
  }

  if connectionPartner == utils.CLIENT_PARTNER {
    id, err := utils.GetIdOfChannel(dataBank, ch)
    if err != nil {
      return err
    }

    responseObj = models.ClientMessage{MessageType: "client_msg", ClientID: id, Payload: fwdMsg.Payload}
    deviceId, err := utils.GetDeviceOfClient(dataBank, id)
    if err != nil {
      return err
    }
    recipient = (*dataBank)[deviceId].Channel

  } else {
    deviceId, err := utils.GetDeviceOfChannel(dataBank, ch)
    if err != nil {
      return err
    }
    responseObj = models.DeviceMessage{MessageType: "device_msg", DeviceID: deviceId, Payload: fwdMsg.Payload}

  }

  response, err := json.Marshal(responseObj)
  if err != nil {
    log.Println("Failed to forward message:", err)
    return err
  }
  
  log.Println("TIME TO FORWARD 2")
  recipient <- models.ChannelMessage{ForwardYesWsNo: true, Message: string(response)}

  return nil
}

