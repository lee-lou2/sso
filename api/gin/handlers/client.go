package handlers

import (
	"github.com/gin-gonic/gin"
	"sso/cmd/client"
	"sso/cmd/user"
	"sso/config/database"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/database/queryset"
	"sso/pkg/security"
	"strconv"
)

// defaultRole 기본 권한명
const defaultRole = "default"

// CreateClientHandler 클라이언트 생성
func CreateClientHandler(c *gin.Context) {
	var payload client.Client

	// 데이터 바인딩
	if err := c.ShouldBindJSON(&payload); err != nil {
		errors.GinErrorJSON(c, errors.New(status.RequestBodyBindingError))
		return
	}
	if payload.Name == "" {
		errors.GinErrorJSON(c, errors.New(status.ClientNameRequired))
		return
	}

	// 데이터베이스 조회
	db, err := database.GetDatabase()
	if err != nil {
		errors.GinErrorJSON(c, err)
		return
	}

	// 존재 여부 확인
	if exists := queryset.Exists(db.Model(&client.Client{}).Where(&client.Client{Name: payload.Name})); exists {
		errors.GinErrorJSON(c, errors.New(status.ExistsClient))
		return
	}

	// 데이터 생성
	db.Create(&payload)
	clientId, err := security.AESCipherEncrypt(strconv.Itoa(int(payload.ID)))
	if err != nil {
		errors.GinErrorJSON(c, errors.New(status.FailedSecretKeyGeneration))
		return
	}
	c.JSON(201, map[string]string{
		"client_id":  clientId,
		"secret_key": payload.SecretKey,
	})
}

// GroupRequestBody 그룹 요청 데이터
type GroupRequestBody struct {
	Name     string `json:"name"`
	ClientId int    `json:"client_id"`
}

// CreateClientGroupHandler 그룹 생성
func CreateClientGroupHandler(c *gin.Context) {
	var payload GroupRequestBody

	if err := c.ShouldBindJSON(&payload); err != nil {
		errors.GinErrorJSON(c, errors.New(status.RequestBodyBindingError))
		return
	}

	// 데이터베이스 조회
	db, err := database.GetDatabase()
	if err != nil {
		errors.GinErrorJSON(c, err)
		return
	}

	// 존재 여부 확인
	if exists := queryset.Exists(db.Model(&user.Group{}).Where(&user.Group{Name: payload.Name})); exists {
		errors.GinErrorJSON(c, errors.New(status.ExistsClient))
		return
	}

	// 그룹 및 권한 생성
	groupObj := user.Group{Name: payload.Name, ClientId: payload.ClientId}
	db.Create(&groupObj)
	db.Create(&user.Role{
		GroupID:     int(groupObj.ID),
		Name:        defaultRole,
		Description: "Default Role",
	})
	c.JSON(201, map[string]interface{}{
		"group_id":     int(groupObj.ID),
		"default_role": defaultRole,
	})
}
