package routers

import (
	"github.com/gin-gonic/gin"
	"sso/api/gin/handlers"
)

// setRoutersV1 라우터 적용 - V1
func setRoutersV1(r *gin.Engine) {
	v1 := r.Group("v1")

	// User
	user := v1.Group("user")
	{
		user.POST("/", handlers.SignUpHandler)
		user.POST("/pending", handlers.PendingUserHandler)
		user.POST("/forgot/password", handlers.ForgotPasswordHandler)
		user.PUT("/password", handlers.SetPasswordHandler)
	}

	// OAuth2
	auth := v1.Group("auth")
	{
		auth.POST("/login", handlers.SignInHandler)
		auth.GET("/:provider/login", handlers.OAuth2RedirectHandler)
		auth.GET("/:provider/callback", handlers.OAuth2CallbackHandler)
	}

	// Client
	client := v1.Group("client")
	{
		client.POST("/", handlers.CreateClientHandler)
		client.POST("/group", handlers.CreateClientGroupHandler)
	}
}
