package initializes

import (
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/global"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/database"
	service "github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/services"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/services/impl"
)

func InitServiceInterface() {
	queries := database.New(global.DB)
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
}
