package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/response"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken, "")
			c.Abort()
			return
		}

		c.Next()
	}
}
