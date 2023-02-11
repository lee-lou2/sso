package user

import (
	"testing"
)

func TestUserSetPassword(t *testing.T) {
	user := User{Email: "test@example.com"}
	user.SetPassword("password")

	if user.Password == "password" {
		t.Errorf("Expected password to be hashed, but got %s", user.Password)
	}
}

func TestGroupBeforeCreate(t *testing.T) {
	group := Group{Name: "Test Group"}
	group.BeforeCreate(nil)

	if group.UUID == "" {
		t.Errorf("Expected UUID to be generated, but got empty string")
	}
}
