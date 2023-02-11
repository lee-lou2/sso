package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"sso/config/database"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/database/queryset"
	"sso/pkg/notify"
)

// CheckPasswordHash 패스워드 확인
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// ForgotPassword 패스워드 찾기
func ForgotPassword(email, state string) error {
	// 1분에 1회만 발송 가능

	// 존재 여부 확인
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}
	if exists := queryset.Exists(db.Model(&User{}).Where(&User{Email: email, Verified: true})); !exists {
		return errors.New(status.NotFoundUser)
	}

	// 인증키 저장
	verifiedCode, err := setVerifiedEmailCode(email)
	if err != nil {
		return err
	}

	// 이메일 작성
	projectName := os.Getenv("PROJECT_NAME")
	serverHost := os.Getenv("PROJECT_API_ENDPOINT")

	// 메세지 선택
	subject := "[" + projectName + "] 패스워드를 잊어버리셨나요?🥲"
	messageFormat := `
	아래 링크 클릭시 간딘히 변경 가능합니다 ☺️<br>
	<a href='%s/password?email=%s&action=forgot&verified_code=%s&state=%s'>패스워드 변경</a>
	`
	message := fmt.Sprintf(messageFormat, serverHost, email, verifiedCode, state)
	// 메일 전송
	if err := notify.SendSMTPEmail(
		email,
		subject,
		message,
	); err != nil {
		return err
	}
	return nil
}

// SetPassword 패스워드 변경
func SetPassword(email, password, verifiedCode string) error {
	var userObj User
	// 인증코드 확인
	if err := verifiedEmailCode(email, verifiedCode); err != nil {
		return err
	}

	// 존재 여부 확인
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}
	db.Where(&User{Email: email, Verified: true}).First(&userObj)
	if userObj.ID != 0 {
		return errors.New(status.NotFoundUser)
	}

	// 패스워드 변경
	userObj.SetPassword(password)
	db.Save(&userObj)
	return nil
}
