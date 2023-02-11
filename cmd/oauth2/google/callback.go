package google

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sso/cmd/client"
	"sso/cmd/oauth2"
	"sso/cmd/user"
	"sso/config/database"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/security"
	"strconv"
)

// OAuth2GoogleCallback 구글 연동/이메일 조회/사용자 생성 및 인증
func OAuth2GoogleCallback(c *gin.Context) {
	var userObj user.User
	var providerObj user.Provider
	var clientObj client.Client
	var groupUserObj user.GroupUser

	code := c.Query("code")
	userId := c.Query("authuser")
	if code == "" || userId == "" {
		errors.GinErrorTemplate(c, errors.New(status.InvalidReturnData))
		return
	}

	db, err := database.GetDatabase()
	if err != nil {
		errors.GinErrorTemplate(c, err)
		return
	}
	// 공급사 존재 여부 확인
	db.Where(&user.Provider{ProviderID: userId, IsActive: true}).Find(&providerObj)
	if providerObj.ID != 0 {
		// 공급사 존재 여부 확인
		db.Where("id = ?", providerObj.UserID).First(&userObj)
		if userObj.ID == 0 {
			errors.GinErrorTemplate(c, errors.New(status.NotFoundUser))
			return
		}
		// 그룹 유저 존재 여부 확인
		db.Where(&user.GroupUser{UserID: int(userObj.ID)}).First(&groupUserObj)
		if groupUserObj.ID == 0 {
			errors.GinErrorTemplate(c, errors.New(status.NotFoundUser))
			return
		}
	} else {
		var group user.Group

		state := c.Query("state")
		db.Where(&user.Group{UUID: state}).First(&group)
		if group.ID == 0 {
			errors.GinErrorTemplate(c, errors.New(status.NotFoundClientGroup))
			return
		}

		// 토큰 발급
		respToken, err := ConfigGoogle().Exchange(c, code)
		if err != nil {
			errors.GinErrorTemplate(c, errors.New(status.FailedGenerateToken))
			return
		}
		// 이메일 조회
		email, err := GetEmail(respToken.AccessToken)
		if err != nil {
			errors.GinErrorTemplate(c, err)
			return
		}

		// 사용자 생성
		createdUserObj, err := user.CreateUser(&user.RequestBody{
			Email:      email,
			ProviderID: userId,
			Provider:   oauth2.ProviderGoogle,
		})
		if err != nil {
			errors.GinErrorTemplate(c, err)
			return
		}
		groupUserObj = user.GroupUser{User: *createdUserObj, GroupId: int(group.ID)}
		db.Create(&groupUserObj)
	}

	// 클라이언트 조회
	db.Where(&client.Client{ID: uint(groupUserObj.GroupId)}).First(&clientObj)
	if clientObj.ID == 0 {
		errors.GinErrorTemplate(c, errors.New(status.NotFoundClientGroup))
		return
	}

	// 사용자 아이디 암호화
	encId, _ := security.AESCipherEncrypt(
		strconv.Itoa(int(userObj.ID)),
		security.CipherConfig{AESCipherKey: clientObj.SecretKey},
	)
	c.Redirect(302, fmt.Sprintf("%s?provider=google&authuser=%s", clientObj.CallbackUri, encId))
}
