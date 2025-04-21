package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.POST("/active")
	}
}
