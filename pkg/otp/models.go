package otp

import (
	"gorm.io/gorm"
	"sso/cmd/user"
)

// OTP OTP 정보
type OTP struct {
	gorm.Model
	UserID    int
	User      user.User `gorm:"constraint:OnDelete:CASCADE;"`
	SecretKey string    `gorm:"not null;unique"`
}
