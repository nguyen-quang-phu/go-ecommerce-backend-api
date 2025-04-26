package account

import (
	"github.com/gin-gonic/gin"
	service "github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/services"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/response"
)

var Login = new(cUserLogin)

type cUserLogin struct{}

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}
