package util

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateKey 암호화키 생성
func GenerateKey(size int) (string, error) {
	ret := make([]byte, size)

	// 키 생성
	if _, err := rand.Read(ret); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ret), nil
}
