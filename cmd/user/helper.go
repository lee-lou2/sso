package user

import (
	"golang.org/x/crypto/bcrypt"
)

// hashPassword 패스워드 해싱
func hashPassword(password string) (string, error) {
	if err := validatePassword(password); err != nil {
		panic(err)
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
