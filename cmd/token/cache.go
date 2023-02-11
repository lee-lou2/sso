package token

import (
	"fmt"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/cache/methods"
)

// SetTokenInCache 토큰 저장
func SetTokenInCache(userId int, tokenType string, sessionId string) error {
	key := fmt.Sprintf("USER_TOKEN__%s__%d", tokenType, userId)
	if err := methods.PushValueSetLimit(key, sessionId, userTokenSetLimit); err != nil {
		return err
	}
	return nil
}

// ExistsTokenInCache 토큰 존재 여부 확인
func ExistsTokenInCache(userId int, tokenType string, sessionId string) error {
	key := fmt.Sprintf("USER_TOKEN__%s__%d", tokenType, userId)
	exists, err := methods.ExistsValue(key, sessionId)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New(status.NotFoundToken)
	}
	return nil
}
