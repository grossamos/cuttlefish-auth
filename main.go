package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grossamos/cuttlefish-auth/handlers"
)

func heathCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
		"about":  "this is a signaling server for cuttlefish",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", heathCheck)
	r.GET("/register_device", handlers.RegisterDevice)

	r.GET("/frontend", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.SetTrustedProxies([]string{})
	r.Run()
}
