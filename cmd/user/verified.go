package user

import (
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/cache/methods"
	"sso/pkg/util"
)

// verifiedCode 인증 코드 확인
func verifiedEmailCode(user, code string) error {
	cacheVerifiedCode, _ := methods.GetValue("VERIFIED_EMAIL_CODE_" + user)
	if cacheVerifiedCode != code {
		return errors.New(status.MismatchedAuthCode)
	}
	return nil
}

func setVerifiedEmailCode(user string) (string, error) {
	// 인증키 저장
	verifiedCode := util.RandStringBytes(8, util.LetterBytes)
	if err := methods.SetValue(
		"VERIFIED_EMAIL_CODE_"+user,
		verifiedCode,
		5*60,
	); err != nil {
		return "", err
	}
	return verifiedCode, nil
}
