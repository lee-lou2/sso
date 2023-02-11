package user

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := hashPassword(password)
	if err != nil {
		t.Errorf("hashPassword returned an error: %v", err)
	}
	if len(hashedPassword) == 0 {
		t.Errorf("hashPassword did not generate a hash")
	}

	// Test that the hashed password is not the same as the original password
	if password == hashedPassword {
		t.Errorf("hashPassword did not hash the password properly, the hash is the same as the original password")
	}
}
