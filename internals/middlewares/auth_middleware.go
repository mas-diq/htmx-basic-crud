package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mas-diq/htmx-basic-crud/internals/utils"
)

// AuthMiddleware is a simple authentication middleware
// This is a placeholder for a real authentication system
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// In a real application, you would check for a valid session or token
		// For this example, we'll just check for a dummy API key
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			// Allow regular browser requests to pass through
			if c.GetHeader("Accept") == "text/html" {
				c.Next()
				return
			}
			utils.ErrorResponse(c, http.StatusUnauthorized, "API key is required")
			c.Abort()
			return
		}

		// In a real app, validate the API key
		if apiKey != "test-api-key" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid API key")
			c.Abort()
			return
		}

		c.Next()
	}
}
