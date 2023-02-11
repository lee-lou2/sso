package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sso/cmd/client"
	"sso/cmd/user"
	"sso/config/database"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/security"
	"strconv"
)

// PasswordHandler 패스워드 설정
func PasswordHandler(c *gin.Context) {
	// 데이터 조회
	action := c.Query("action")
	state := c.Query("state")
	email := c.Query("email")
	verifiedCode := c.Query("verified_code")

	c.HTML(http.StatusOK, "password.html", gin.H{
		"action":       action,
		"email":        email,
		"verifiedCode": verifiedCode,
		"state":        state,
	})
}

// EmailHandler 회원가입/패스워드 설정
func EmailHandler(c *gin.Context) {
	var groupObj user.Group
	var clientObj client.Client

	// 데이터 조회
	action := c.Query("action")
	state := c.Query("state")

	if action != "signup" && action != "forgot" {
		errors.GinErrorTemplate(c, errors.New(status.PageStatusNotSet))
		return
	}

	// 코드 확인
	db, err := database.GetDatabase()
	if err != nil {
		errors.GinErrorTemplate(c, err)
		return
	}

	db.Where(&user.Group{UUID: state}).First(&groupObj)
	if groupObj.ID == 0 {
		errors.GinErrorTemplate(c, errors.New(status.NotFoundClientGroup))
		return
	}
	db.Where("id = ?", groupObj.ClientId).First(&clientObj)
	if clientObj.ID == 0 {
		errors.GinErrorTemplate(c, errors.New(status.NotFoundClient))
		return
	}

	c.HTML(http.StatusOK, "email.html", gin.H{
		"isCreate": action == "signup",
		"state":    state,
		"homepage": clientObj.Homepage,
		"action":   action,
	})
}

// clientSetResponse 클라이언트 세트 반환 값
type clientSetResponse struct {
	Client client.Client
	Group  user.Group
}

// getClient 클라이언트/그룹 조회
func getClient(code, group string) (*clientSetResponse, error) {
	var clientObj client.Client

	// 1. 코드 확인
	idStr, err := security.AESCipherDecrypt(code)
	if err != nil {
		return nil, errors.New(status.InvalidClientCode)
	}
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	clientId, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, errors.New(status.InvalidClientCode)
	}
	db.Where("id = ?", clientId).First(&clientObj)
	if clientObj.ID == 0 {
		return nil, errors.New(status.NotFoundClient)
	}

	// 2. 권한 확인
	var groupObj user.Group

	groupName, err := security.AESCipherDecrypt(
		group,
		security.CipherConfig{AESCipherKey: clientObj.SecretKey},
	)
	if err != nil {
		return nil, errors.New(status.InvalidClientCode)
	}

	// 존재하는 그룹인지 확인
	db.Where(&user.Group{Name: groupName, Client: clientObj}).First(&groupObj)
	if groupObj.ID == 0 {
		return nil, errors.New(status.NotFoundClientGroup)
	}
	return &clientSetResponse{Group: groupObj, Client: clientObj}, nil
}

// LoginHandler 로그인 페이지
func LoginHandler(c *gin.Context) {
	code := c.Query("code")
	group := c.Query("group")

	// 클라이언트 정보 조회
	clientSet, err := getClient(code, group)
	if err != nil {
		errors.GinErrorTemplate(c, errors.New(status.InvalidClientCode))
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"state": clientSet.Group.UUID,
		"code":  code,
		"group": group,
	})
}
