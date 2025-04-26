package impl

import (
	"context"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/database"
	service "github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/services"
)

type sUserLogin struct {
	r *database.Queries
}

// GeneratePasswordRegister implements service.IUserLogin.
func (s *sUserLogin) GeneratePasswordRegister(ctx context.Context) error {
	panic("unimplemented")
}

// Login implements service.IUserLogin.
func (s *sUserLogin) Login(ctx context.Context) error {
	panic("unimplemented")
}

// Register implements service.IUserLogin.
func (s *sUserLogin) Register(ctx context.Context) error {
	panic("unimplemented")
}

// VerrifyOTP implements service.IUserLogin.
func (s *sUserLogin) VerrifyOTP(ctx context.Context) error {
	panic("unimplemented")
}

func NewUserLoginImpl(r *database.Queries) service.IUserLogin {
	return &sUserLogin{r: r}
}
