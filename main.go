package main

import (
	"github.com/gin-gonic/gin"
)

func heathCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
		"about":  "this is a signaling server for cuttlefish",
	})
}

func main() {
	r := gin.Default()
	r.StaticFile("/", "./webui/index.html")
	r.StaticFile("/index.css", "./webui/index.css")
	r.StaticFile("/style.css", "./webui/style.css")
  r.Static("/js/", "./webui/js/")
	r.GET("/register_device", WSHandler)
	r.GET("/connect_client", WSHandler)
	r.GET("/frontend", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.SetTrustedProxies([]string{})
  r.RunTLS("0.0.0.0:8443", "/tmp/ca/certificate.pem", "/tmp/ca/privatekey.pem")
  // r.Run("0.0.0.0:7443")
}
