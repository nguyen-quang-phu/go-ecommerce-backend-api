package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003

	ErrInvalidToken      = 30001
	ErrInvalidOTP      = 30002

	ErrCodeUserHasExists = 50001
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvalidOTP:      "OTP is invalid",
	ErrCodeUserHasExists: "User has already registered",
}
