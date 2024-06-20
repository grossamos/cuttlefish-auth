package auth

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const authFile = "./auth.json"

func GetGinBasicAuth() gin.HandlerFunc {
  authInfo := getAuthInformation()
	return gin.BasicAuth(authInfo.Users)
}

func getAuthInformation() AuthFileStructure {
	authInfoRaw, err := os.ReadFile(authFile)
	if err != nil {
		log.Fatal("Couldn't read users file, aborting...")
	}

  var authInfo AuthFileStructure
	err = json.Unmarshal(authInfoRaw, &authInfo)
	if err != nil {
		log.Fatal("Couldn't unmarshal users json, aborting...")
	}

  return authInfo
}

type AuthFileStructure struct {
	Users  map[string]string `json:"users"`
	JwtKey string            `json:"jwt-key"`
}
