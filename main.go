package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grossamos/cuttlefish-auth/auth"
	"github.com/grossamos/cuttlefish-auth/endpoints"
	"github.com/grossamos/cuttlefish-auth/models"
	"github.com/grossamos/cuttlefish-auth/utils"
)

func heathCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
		"about":  "this is a signaling server for cuttlefish",
	})
}

func wrapperFunction(deviceRegistry *map[string]models.DataBankEntry, communicationPartner string) gin.HandlerFunc {
  return func (c *gin.Context) {
    endpoints.WSHandler(c, deviceRegistry, communicationPartner)
  }
}
func wrapperFunctionDevices(deviceRegistry *map[string]models.DataBankEntry) gin.HandlerFunc {
  return func (c *gin.Context) {
    endpoints.DeviceHandler(c, deviceRegistry)
  }
}

func main() {
  jwtState := auth.InitializeJwtState()
  deviceRegistry := make(map[string]models.DataBankEntry)
  WSHandlerWrapperDevice := wrapperFunction(&deviceRegistry, utils.DEVICE_PARTNER)
  WSHandlerWrapperClient := wrapperFunction(&deviceRegistry, utils.CLIENT_PARTNER)
  deviceHandkerWraooer := wrapperFunctionDevices(&deviceRegistry)
  authMiddleware := auth.AuthMiddleware(&jwtState)
  loginHandler := auth.GetLoginHandler(jwtState)


	r := gin.Default()
	r.StaticFile("/", "./webui/index.html")
	r.StaticFile("/index.css", "./webui/index.css")
	r.StaticFile("/login.css", "./webui/login.css")
	r.StaticFile("/login.html", "./webui/login.html")
	r.StaticFile("/client.html", "./webui/client.html")
	r.StaticFile("/controls.css", "./webui/controls.css")
	r.StaticFile("/style.css", "./webui/style.css")
  r.Static("/js/", "./webui/js/")
  r.POST("/login", auth.GetGinBasicAuth(), loginHandler)
  r.GET("/register_device", WSHandlerWrapperDevice)
  r.GET("/connect_client", authMiddleware, WSHandlerWrapperClient)
  r.GET("/devices", authMiddleware, deviceHandkerWraooer)

	r.SetTrustedProxies([]string{})
  r.RunTLS("0.0.0.0:8443", "./certs/server.crt", "./certs/server.key")
  // r.Run("0.0.0.0:7443")
}
