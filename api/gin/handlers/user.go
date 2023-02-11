package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sso/cmd/user"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/security"
	"strconv"
)

// pendingRequestBody 대기 사용자 요청 데이터
type pendingRequestBody struct {
	Email string `json:"email"`
	State string `json:"state"`
}

// PendingUserHandler 대기 사용자 처리
func PendingUserHandler(c *gin.Context) {
	var payload pendingRequestBody

	// 데이터 바인딩
	if err := c.ShouldBindJSON(&payload); err != nil {
		errors.GinErrorJSON(c, errors.New(status.RequestBodyBindingError))
		return
	}

	// 대기 사용자 이메일 전송
	if err := user.PendingUser(payload.Email, payload.State); err != nil {
		errors.GinErrorJSON(c, err)
		return
	}
	c.JSON(200, map[string]bool{"is_completed": true})
}

// signUpRequestBody 요청 데이터
type signUpRequestBody struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Code         string `json:"code"`
	VerifiedCode string `json:"verified_code"`
}

// SignUpHandler 회원 가입
func SignUpHandler(c *gin.Context) {
	var payload signUpRequestBody

	// 데이터 바인딩
	if err := c.ShouldBindJSON(&payload); err != nil {
		errors.GinErrorJSON(c, errors.New(status.RequestBodyBindingError))
		return
	}

	// 클라이언트 및 그룹 이상여부 확인
	clientSet, err := user.ValidateClient(payload.Code)
	if err != nil {
		errors.GinErrorJSON(c, errors.New(status.NotFoundClient))
		return
	}

	// 사용자 생성
	if err := user.CreateDefaultUser(
		payload.Email,
		payload.Password,
		payload.VerifiedCode,
	); err != nil {
		errors.GinErrorJSON(c, err)
		return
	}

	// 로그인 페이지로 이동
	c.JSON(201, map[string]string{"code": clientSet.ClientId, "group": clientSet.GroupName})
}

// signInRequestBody 요청 데이터
type signInRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

// SignInHandler 자체 사용자 로그인 핸들러
func SignInHandler(c *gin.Context) {
	var payload signInRequestBody

	// 데이터 바인딩
	if err := c.ShouldBindJSON(&payload); err != nil {
		errors.GinErrorJSON(c, errors.New(status.RequestBodyBindingError))
		return
	}

	// 클라이언트 및 그룹 이상여부 확인
	clientObj, err := user.SelectClientSet(payload.Code)
	if err != nil {
		errors.GinErrorJSON(c, errors.New(status.NotFoundClient))
		return
	}

	userId, err := user.SignIn(payload.Email, payload.Password)
	if err != nil {
		errors.GinErrorJSON(c, err)
		return
	}

	// 사용자 아이디 암호화
	encId, err := security.AESCipherEncrypt(
		strconv.Itoa(userId),
		security.CipherConfig{AESCipherKey: clientObj.Client.SecretKey},
	)
	if err != nil {
		errors.GinErrorJSON(c, err)
		return
	}
	c.JSON(200, map[string]string{
		"callback_uri": fmt.Sprintf("%s?provider=default&authuser=%s", clientObj.Client.CallbackUri, encId),
	})
}

// ForgotPasswordHandler 패스워드 찾기 핸들러
func ForgotPasswordHandler(c *gin.Context) {
	var payload pendingRequestBody

	// 데이터 바인딩
	if err := c.ShouldBindJSON(&payload); err != nil {
		errors.GinErrorJSON(c, errors.New(status.RequestBodyBindingError))
		return
	}

	// 대기 사용자 이메일 전송
	if err := user.ForgotPassword(payload.Email, payload.State); err != nil {
		errors.GinErrorJSON(c, err)
		return
	}
	c.JSON(200, map[string]bool{"is_completed": true})
}

// SetPasswordHandler 패스워드 변경 핸들러
func SetPasswordHandler(c *gin.Context) {
	var payload signUpRequestBody

	// 데이터 바인딩
	if err := c.ShouldBindJSON(&payload); err != nil {
		errors.GinErrorJSON(c, errors.New(status.RequestBodyBindingError))
		return
	}

	// 클라이언트 및 그룹 이상여부 확인
	clientSet, err := user.ValidateClient(payload.Code)
	if err != nil {
		errors.GinErrorJSON(c, errors.New(status.NotFoundClient))
		return
	}

	// 사용자 생성
	if err := user.SetPassword(
		payload.Email,
		payload.Password,
		payload.VerifiedCode,
	); err != nil {
		errors.GinErrorJSON(c, err)
		return
	}

	// 로그인 페이지로 이동
	c.JSON(200, map[string]string{"code": clientSet.ClientId, "group": clientSet.GroupName})
}
