package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				fmt.Printf("Panic occurred: %v\n", err)

				// Return a unified error response
				c.JSON(500, gin.H{
					"code":    500,
					"message": "Internal Server Error",
					"error":   fmt.Sprintf("%v", err),
				})
				c.Abort() // Stop further execution
			}
		}()
		c.Next()
	}
}
