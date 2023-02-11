package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
	"sso/config/errors"
	"sso/config/errors/status"
)

// CipherConfig 정보
type CipherConfig struct {
	AESCipherKey string
}

// aesValidate 데이터 유효성 검사
func aesValidate(data string) error {
	if data == "" {
		return errors.New(status.ConversionTextNotFound)
	}
	return nil
}

// AESCipherEncrypt AES 암호화
func AESCipherEncrypt(data string, config ...CipherConfig) (string, error) {
	var key []byte
	if err := aesValidate(data); err != nil {
		return "", err
	}
	if len(config) > 0 && config[0].AESCipherKey != "" {
		key = []byte(config[0].AESCipherKey)
	} else {
		key = []byte(os.Getenv("AES_CIPHER_KEY"))
	}
	plainText := []byte(data)

	// AES 대칭키 암호화 블록 생성
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 암호화
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	encText := base64.URLEncoding.EncodeToString(cipherText)
	return encText, nil
}

// AESCipherDecrypt AES 복호화
func AESCipherDecrypt(encText string, config ...CipherConfig) (string, error) {
	var key []byte

	if err := aesValidate(encText); err != nil {
		return "", err
	}
	if len(config) > 0 && config[0].AESCipherKey != "" {
		key = []byte(config[0].AESCipherKey)
	} else {
		key = []byte(os.Getenv("AES_CIPHER_KEY"))
	}

	cipherText, err := base64.URLEncoding.DecodeString(encText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New(status.ShortCipherBlockSize)
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
