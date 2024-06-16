package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/grossamos/cuttlefish-auth/models"
)

func WSHandler(ctx *gin.Context, dataBank *map[string]models.DataBankEntry, communicationPartner string)  {
  conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
  }

  defer conn.Close()

  ch := make(chan models.ChannelMessage)
  go channelIncomingMessages(conn, ch)

  for chmsg := range ch {
    if chmsg.ForwardYesWsNo {
      log.Println("Yayyy, forwarding")
      conn.WriteMessage(websocket.TextMessage, []byte(chmsg.Message))
      continue
    } else {
      message := chmsg.Message
      responses, err := handleIncomingMessage(message, dataBank, communicationPartner, ch)
      if err != nil {
        response := "{\"error\": \"" + err.Error() + "\" }"
        err = conn.WriteMessage(websocket.TextMessage, []byte(response))
        if err != nil {
          log.Println("Error writing response:", err)
          break
        }
        continue
      }

      for _, response := range responses {
        err = conn.WriteMessage(websocket.TextMessage, []byte(response))
        if err != nil {
          log.Println("Error writing response:", err)
          break
        }
      }
    }
  }
}


type mesageTypeContainer struct {
  MessageType string `json:"message_type"`
}


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func channelIncomingMessages(conn *websocket.Conn, ch chan models.ChannelMessage) error {
  for {
    _, message, err := conn.ReadMessage()
    if err != nil {
      log.Println("Read error:", err)
      return err
    }
    log.Println(string(message))
    ch <- models.ChannelMessage{ForwardYesWsNo: false, Message: string(message)}
  }
}


func handleIncomingMessage(message string, dataBank *map[string]models.DataBankEntry, connectionPartner string, ch chan models.ChannelMessage) ([]string, error) {
  var msgTypeContainer mesageTypeContainer
  response := make([]string, 0)
  var err error

  err = json.Unmarshal([]byte(message), &msgTypeContainer)
  if err != nil {
    log.Println("Error unmarshaling incoming json: ", err)
    return nil, err
  }

  if msgTypeContainer.MessageType == REGISTER_MESSAGE_TYPE || msgTypeContainer.MessageType == CONNECT_MESSAGE_TYPE {
    responseMessage, err := HandleRegistration(message, dataBank, connectionPartner, ch)
    if err != nil {
      return nil, err
    }
    response = append(response, responseMessage)

    if msgTypeContainer.MessageType == CONNECT_MESSAGE_TYPE {
      log.Println(CONNECT_MESSAGE_TYPE)
      responseMessage, err := HandleDeviceInfoPropagation(message, dataBank)
      if err != nil {
        return nil, err
      }
      response = append(response, responseMessage)
    }
  } else if msgTypeContainer.MessageType == FORWARD_MESSAGE_TYPE {
    err = ForwardMessageFromSender(dataBank, message, connectionPartner, ch)
    if err != nil {
      return nil, err
    }
  } else {
    log.Println("Unkown message_type: ", msgTypeContainer.MessageType)
    return nil, err
  }

  return response, nil
}

func DeviceHandler(c *gin.Context, dataBank *map[string]models.DataBankEntry) {
  devices := make([]string, 0)
  for deviceId, _ := range *dataBank {
    devices = append(devices, deviceId)
  }
  c.JSON(200, devices)
}
