package user

import (
	"fmt"
	"sso/cmd/client"
	"sso/config/database"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/security"
)

// validatePassword 패스워드 유효성 검사
func validatePassword(password string) error {
	if password == "" {
		return nil
	}
	return nil
}

// ClientSetResponse 클라이언트 반환 값
type ClientSetResponse struct {
	Client client.Client
	Group  Group
}

// SelectClientSet 클라이언트 데이터 조회
func SelectClientSet(uuid string) (*ClientSetResponse, error) {
	var groupObj Group
	var clientObj client.Client

	// 그룹 조회
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	db.Where(&Group{UUID: uuid}).First(&groupObj)
	if groupObj.ID == 0 {
		return nil, errors.New(status.NotFoundClientGroup)
	}

	// 클라이언트 조회
	db.Where("id = ?", groupObj.ClientId).First(&clientObj)
	if clientObj.ID == 0 {
		return nil, errors.New(status.NotFoundClient)
	}
	return &ClientSetResponse{Client: clientObj, Group: groupObj}, nil
}

// ValidateClientResponse 클라이언트 반환 값
type ValidateClientResponse struct {
	ClientId  string
	GroupName string
}

// ValidateClient 클라이언트 이상 여부 확인
func ValidateClient(uuid string) (*ValidateClientResponse, error) {
	clientSet, err := SelectClientSet(uuid)
	if err != nil {
		return nil, err
	}

	// 암호화 데이터 조회
	encGroupId, err := security.AESCipherEncrypt(fmt.Sprintf("%d", clientSet.Group.ID))
	if err != nil {
		return nil, err
	}
	encGroupName, err := security.AESCipherEncrypt(clientSet.Group.Name, security.CipherConfig{
		AESCipherKey: clientSet.Client.SecretKey,
	})
	return &ValidateClientResponse{ClientId: encGroupId, GroupName: encGroupName}, nil
}
