package convert

import (
	"encoding/base64"
)

// Base64Encode 인코딩
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64Decode 디코딩
func Base64Decode(encValue string) (string, error) {
	value, err := base64.StdEncoding.DecodeString(encValue)
	return string(value), err
}
