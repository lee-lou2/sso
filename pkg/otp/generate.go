package otp

import (
	"github.com/xlzd/gotp"
	"sso/config/database"
	"sso/pkg/database/queryset"
)

// GetSecretKey 시크릿키 조회
func GetSecretKey(userId int) (string, error) {
	var otp OTP

	// 저장된 정보 조회
	db, err := database.GetDatabase()
	if err != nil {
		return "", err
	}
	db.Where(&OTP{UserID: userId}).First(&otp)
	if otp.ID == 0 {
		return "", nil
	}
	return otp.SecretKey, nil
}

// CreateSecretKey OTP 시크릿키 생성
func CreateSecretKey(userId int) (*OTP, error) {
	var otp OTP
	var secretKey string

	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	// 고유한 시크릿키 생성
	for {
		secretKey = gotp.RandomSecret(otpSecretLength)
		if exists := queryset.Exists(db.Model(&OTP{}).Where(&OTP{SecretKey: secretKey})); !exists {
			break
		}
	}
	otp = OTP{
		UserID:    userId,
		SecretKey: secretKey,
	}
	db.Create(&otp)
	return &otp, nil
}

// GetOTPUri OTP Uri 조회
func GetOTPUri(userId int) (string, error) {
	// 시크릿키 조회
	secretKey, err := GetSecretKey(userId)
	if err != nil {
		return "", err
	}
	// 데이터가 존재하지 않는 경우 신규 생성
	if secretKey == "" {
		otp, err := CreateSecretKey(userId)
		if err != nil {
			return "", err
		}
		secretKey = otp.SecretKey
	}
	// Uri 생성
	return gotp.NewDefaultTOTP(secretKey).ProvisioningUri(otpAccountName, otpIssuerName), nil
}
