package initializes

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/middlewares"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/response"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenticateMiddleware())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"pong"})
}
