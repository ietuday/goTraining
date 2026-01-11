package middlewares

import (
	"net/http"

	"example.com/rest-api/auth"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		rawToken, err := auth.ExtractBearerToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
			c.Abort()
			return
		}

		claims, err := auth.VerifyToken(rawToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
			c.Abort()
			return
		}

		// Optional: make userId/email available to handlers
		if v, ok := claims["userId"].(float64); ok {
			c.Set("userId", int64(v))
		}
		if email, ok := claims["email"].(string); ok {
			c.Set("email", email)
		}

		c.Next()
	}
}
	