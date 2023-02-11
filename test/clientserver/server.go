package main

import (
	"github.com/gin-gonic/gin"
	"sso/config/env"
)

const ServerHost = "http://localhost:3000"

func main() {
	// 환경 설정
	env.Load()

	// 마이그레이션
	AutoMigrate()

	// 기본 엔진 생성
	r := gin.Default()
	r.LoadHTMLGlob("./web/*")

	// 라우터
	r.GET("/", testMainPageHandler)
	r.POST("/client", testCreateClientHandler)
	r.POST("/client/group", testCreateGroupHandler)
	r.GET("/login", testLoginHandler)
	r.GET("/callback", testCallbackHandler)
	r.GET("/status", testStatusCheckHandler)
	r.GET("/logout", testLogoutHandler)

	r.Run(":8081")
}
