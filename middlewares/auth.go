package middlewares

import (
	"btpn-go/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware validates JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.JSONError(c, http.StatusUnauthorized, "Authorization header is missing")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			helpers.JSONError(c, http.StatusUnauthorized, "Bearer token is missing")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret_key"), nil // Replace with your secret key
		})

		if err != nil {
			helpers.JSONError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", uint(claims["user_id"].(float64)))
			c.Next()
		} else {
			helpers.JSONError(c, http.StatusUnauthorized, "Invalid token claims")
			c.Abort()
		}
	}
}
