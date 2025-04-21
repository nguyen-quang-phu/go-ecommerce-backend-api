package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	service "github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/services"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/vo"
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
	var params vo.UserRegistratorRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
		return
	}

	fmt.Printf("Email params: %s", params.Email)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
