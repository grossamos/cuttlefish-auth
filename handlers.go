package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WSHandler(ctx *gin.Context)  {
  conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
  }

  defer conn.Close()

  for {
    _, message, err := conn.ReadMessage()
    if err != nil {
      log.Println("Read error:", err)
      break
    }

    response, err := HandleSignaling(DEVICE_PARTNER, string(message))
    if err != nil {
      break
    }

    err = conn.WriteMessage(websocket.TextMessage, []byte(response))
    if err != nil {
      log.Println("Error writing response:", err)
      break
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

