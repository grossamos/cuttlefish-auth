package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


func RegisterDevice(ctx *gin.Context)  {
  var msgContainer mesageTypeContainer
  
  conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
  }

  for {
    _, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

    err = json.Unmarshal(message, &msgContainer)
		if err != nil {
      log.Println("Unable to unmarshal json:", err)
			break
		}

    fmt.Println(msgContainer.MessageType, message)
    break

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

