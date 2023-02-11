package otp

import (
	"github.com/xlzd/gotp"
	"sso/config/errors"
	"sso/config/errors/status"
)

// ValidateOTP OTP 번호 확인
func ValidateOTP(userId int, otpNumber string) (bool, error) {
	secretKey, err := GetSecretKey(userId)
	if err != nil {
		return false, err
	}
	if secretKey == "" {
		return false, errors.New(status.NotFoundOTPSecretKey)
	}
	return gotp.NewDefaultTOTP(secretKey).Now() == otpNumber, nil
}
