package routers

import (
	"github.com/gin-gonic/gin"
	"sso/api/gin/handlers"
)

// setRoutersWeb 웹 페이지 라우터
func setRoutersWeb(r *gin.Engine) {
	r.GET("/", handlers.LoginHandler)
	r.GET("/email", handlers.EmailHandler)
	r.GET("/password", handlers.PasswordHandler)
}
