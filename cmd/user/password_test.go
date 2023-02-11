package user

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err := CheckPasswordHash(password, string(hashedPassword)); err != nil {
		t.Errorf("CheckPasswordHash returned an error: %v", err)
	}

	if err := CheckPasswordHash("wrong_password", string(hashedPassword)); err == nil {
		t.Errorf("CheckPasswordHash should have returned an error but did not")
	}
}
