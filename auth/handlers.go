package auth

import (
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetLoginHandler(jwtState JwtState) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    authHeader := ctx.GetHeader("Authorization")

    if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication header is required"})
        ctx.Abort()
        return
    }

    encodedCredentials := authHeader[6:]

    decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to decode authentication header"})
        ctx.Abort()
        return
    }

    decodedCredentials := string(decodedBytes)

    credentials := strings.SplitN(decodedCredentials, ":", 2)
    if len(credentials) != 2 {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid basic auth format"})
        ctx.Abort()
        return
    }

    token, err := GenerateNewToken(&jwtState, credentials[0])
    if err != nil {
      fmt.Println(err.Error())
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
        ctx.Abort()
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": token})
  }
}

func OnlyLocalhost() gin.HandlerFunc {
    return func(c *gin.Context) {
        clientIP := c.ClientIP()
        if !isLocalhost(clientIP) {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
            return
        }
        c.Next()
    }
}

func isLocalhost(ip string) bool {
    if strings.Contains(ip, ":") {
        ip, _, _ = net.SplitHostPort(ip)
    }
    return ip == "127.0.0.1" || ip == "::1"
}
