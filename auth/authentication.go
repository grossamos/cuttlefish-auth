package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func InitializeJwtState() JwtState {
  jwtKey := getAuthInformation().JwtKey
  return JwtState{
    jwtKey: []byte(jwtKey),
  }
}

func GenerateNewToken(jwtState *JwtState, username string) (string, error) {
  claims := jwt.MapClaims{
        "username": username, // kept in for completeness, but not used
        "exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(jwtState.jwtKey)
    if err != nil {
        return "", err
    }
  return tokenString, nil
}

func AuthMiddleware(jtwState *JwtState) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    tokenString := ctx.Query("token")
    if len(tokenString) == 0 {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
        ctx.Abort()
        return
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      // Validate the signing method
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
      }
      return jtwState.jwtKey, nil
    })

    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        ctx.Abort()
        return
    }

    // Validate the token claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      if exp, ok := claims["exp"].(float64); ok {
          if time.Unix(int64(exp), 0).Before(time.Now()) {
              ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
              ctx.Abort()
              return
          }
      } else {
          ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token has no expiration"})
          ctx.Abort()
          return
      }
    } else {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        ctx.Abort()
        return
    }

    ctx.Next()
  }
}

type JwtState struct {
  jwtKey []byte
}
