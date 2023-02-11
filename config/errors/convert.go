package errors

import (
	"github.com/gin-gonic/gin"
)

// GinErrorJSON Gin Error 로 변환
func GinErrorJSON(c *gin.Context, err error) {
	_, statusCode, message := ErrorDetail(err)
	c.AbortWithStatusJSON(statusCode, &gin.H{"message": message})
}

// GinErrorTemplate Gin Error 페이지로 이동
func GinErrorTemplate(c *gin.Context, err error) {
	errorCode, statusCode, message := ErrorDetail(err)
	c.HTML(statusCode, "error.html", gin.H{
		"message":     message,
		"status_code": statusCode,
		"error_code":  errorCode,
	})
}

// ErrorDetail 오류 정보
func ErrorDetail(err error) (int, int, string) {
	obj, ok := err.(*Error)
	if !ok {
		return 0, 500, err.Error()
	}
	return obj.Code, obj.StatusCode, obj.Message
}
