package random

import (
	"math/rand"
	"time"
)

func GenerateSixDigitOtp() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 1e5 + rng.Intn(9e5)
	return otp
}
