package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/forumapp/internal/configs"
	"github.com/rdy24/forumapp/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT

	return func(c *gin.Context) {
		// Get the token from the header
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)

		// Check if the token is empty
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("authorization token is required"))
			return
		}

		userId, username, err := jwt.ValidateToken(header, secretKey)

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT

	return func(c *gin.Context) {
		// Get the token from the header
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)

		// Check if the token is empty
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("authorization token is required"))
			return
		}

		userId, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}
