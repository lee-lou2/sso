package user

import (
	"fmt"
	"os"
	"sso/cmd/oauth2"
	"sso/config/database"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/database/queryset"
	"sso/pkg/notify"
)

// RequestBody 요청 데이터
type RequestBody struct {
	Email      string
	Password   string
	ProviderID string
	Provider   int
	IsVerified bool
}

// PendingUser 가입 대기 사용자
func PendingUser(email, state string) error {
	// 1분에 1회만 발송 가능

	// 존재 여부 확인
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}
	if exists := queryset.Exists(db.Model(&User{}).Where(&User{Email: email})); exists {
		return errors.New(status.UserAlreadyExists)
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
	subject := "[" + projectName + "] 회원 가입을 진행하시겠습니까?🤩"
	messageFormat := `
	아래 링크 클릭시 간딘히 회원 가입이 완료됩니다 ☺️<br>
	<a href='%s/password?email=%s&action=signup&verified_code=%s&state=%s'>회원 가입</a>
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

// CreateDefaultUser 기본 사용자 생성
func CreateDefaultUser(email, password, verifiedCode string) error {
	// 인증코드 확인
	if err := verifiedEmailCode(email, verifiedCode); err != nil {
		return err
	}

	// 존재 여부 확인
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}
	if exists := queryset.Exists(db.Model(&User{}).Where(&User{Email: email})); exists {
		return errors.New(status.UserAlreadyExists)
	}

	// 사용자 생성
	_, err = CreateUser(&RequestBody{
		Email:      email,
		Password:   password,
		IsVerified: true,
		Provider:   oauth2.ProviderDefault,
	})
	if err != nil {
		return err
	}
	return nil
}

// CreateUser 사용자 생성
func CreateUser(req *RequestBody) (*User, error) {
	user := User{
		Email:    req.Email,
		Verified: req.IsVerified,
	}

	// 기본 가입자의 경우 패스워드 지정
	if req.Provider == oauth2.ProviderDefault {
		user.SetPassword(req.Password)
	}
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	// 공급사 정보 저장
	provider := Provider{
		User:     user,
		Provider: req.Provider,
	}
	if err := db.Create(&provider).Error; err != nil {
		return nil, err
	}

	// Oauth2 를 통한 회원 가입
	if req.Provider != oauth2.ProviderDefault {
		// 인증 완료
		user.Verified = true
		if err := db.Save(&user).Error; err != nil {
			return nil, err
		}

		// 공급사 사용자 정보 저장
		provider.ProviderID = req.ProviderID
		if err := db.Save(&provider).Error; err != nil {
			return nil, err
		}
	}
	return &user, nil
}
