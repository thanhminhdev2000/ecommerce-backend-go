package middlewares

import (
	"github.com/gin-gonic/gin"
	response "github.com/thanhminhdev2000/ecommerce-backend-go/pkg/repsponse"
)

func AuthenMiddleware() gin.HandlerFunc {
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