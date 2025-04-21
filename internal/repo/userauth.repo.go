package repo

import (
	"fmt"
	"time"

	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/global"
)

type IUserAuthRepository interface {
	AddOtp(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

// AddOtp implements IUserAuthRepository.
func (u *userAuthRepository) AddOtp(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("urs:%s:otp", email)
	fmt.Println("key:", key)
	return global.Rdb.SetEx(
		ctx,
		key,
		otp,
		time.Duration(time.Duration(expirationTime)),
	).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
