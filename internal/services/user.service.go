package service

import (
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/repo"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
