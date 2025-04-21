package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user", userController.Register)
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}
	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/otp")
	}
}
