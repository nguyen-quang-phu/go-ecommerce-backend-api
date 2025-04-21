package controllers

import (
	"github.com/gin-gonic/gin"
	service "github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/services"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}
