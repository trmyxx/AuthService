package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token is empty"})
			return
		}

		// тут проверка JWT
		userID, err := validateToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			return
		}

		// кладём данные в контекст
		c.Set("user_id", userID)

		c.Next()
	}
}

func validateToken(token string) (string, error) {
	auths := strings.Split(token, " ")
	if len(auths) == 2 && auths[0] == "Bearer" {
		userID := auths[1]
		return userID, nil
	}
	return "", errors.New("invalid token")
}
