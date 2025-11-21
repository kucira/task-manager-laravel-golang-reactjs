package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"task-manager-fullstack/models"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			err := fmt.Errorf("Unauthorized")
			models.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", err)
			c.Abort()
			return
		}

		// Expecting format: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// tokenString := parts[1]

		// Put user ID in context (optional)
		// c.Set("userId", claims.UserID)

		c.Next()
	}
}
