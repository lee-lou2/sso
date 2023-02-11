package errors

import (
	"fmt"
)

// Error 오류 스키마
type Error struct {
	Code       int
	StatusCode int
	Message    string
}

// Error 오류 메세지
func (e *Error) Error() string {
	return fmt.Sprintf("%d, [%d] %s", e.Code, e.StatusCode, e.Message)
}

// New 오류 생성
func New(codeGroup []interface{}, defaultErrors ...error) error {
	errorCode := codeGroup[0].(int)
	statusCode := codeGroup[1].(int)
	message := codeGroup[2].(string)
	err := &Error{
		Message:    message,
		StatusCode: statusCode,
		Code:       errorCode,
	}

	// 오류 알람
	Alert(err)

	// 기본 오류들 알람
	for _, _err := range defaultErrors {
		Alert(_err)
	}
	return err
}
