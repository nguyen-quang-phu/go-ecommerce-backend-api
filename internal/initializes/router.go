package initializes

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/global"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/router"
)

func setMode() *gin.Engine {
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		return gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		return gin.New()
	}
}

func InitRouter() *gin.Engine {
	r := setMode()
	manageRouter := router.RouterGroupApp.Manage
	userRouter := router.RouterGroupApp.User
	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	{
		manageRouter.InitUserRouter(mainGroup)
		manageRouter.InitAdminRouter(mainGroup)
	}
	return r
}
