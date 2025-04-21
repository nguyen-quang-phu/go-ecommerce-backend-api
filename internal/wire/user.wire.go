//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/controllers"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/repo"
	service "github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/services"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controllers.NewUserController,
	)

	return new(controllers.UserController), nil
}
