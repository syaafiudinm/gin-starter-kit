package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC RECOVERED] %v\n%s", err, debug.Stack())

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Internal server error",
					"error":   "An unexpected error occurred. Please try again later.",
				})
			}
		}()

		c.Next()
	}
}
