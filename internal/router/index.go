package router

import (
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/router/manage"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/router/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
