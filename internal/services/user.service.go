package service

import (
	"fmt"
	"time"

	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/repo"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/ultils/crypto"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/ultils/random"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	hashEmail := crypto.GetHash(email)
	fmt.Printf("ashEmail: %s\n", hashEmail)
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST" {
		otp = 123456
	}
	fmt.Printf("Otp is : %d\n", otp)

	err := us.userAuthRepo.AddOtp(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	return response.ErrCodeSuccess
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}
