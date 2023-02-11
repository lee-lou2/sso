package user

import (
	"sso/config/database"
	"sso/config/errors"
	"sso/config/errors/status"
)

// SignIn 로그인
func SignIn(email, password string) (int, error) {
	var userObj User

	db, err := database.GetDatabase()
	if err != nil {
		return 0, err
	}
	db.Where(User{Email: email}).First(&userObj)
	if userObj.ID == 0 {
		return 0, errors.New(status.InvalidUserInfo)
	}

	// 패스워드 확인
	if err := CheckPasswordHash(password, userObj.Password); err != nil {
		return 0, errors.New(status.InvalidUserInfo)
	}
	return int(userObj.ID), nil
}
