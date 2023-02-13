package gin

import (
	"github.com/gin-gonic/gin"
	"log"
	"sso/api/gin/middlewares"
	"sso/api/gin/routers"
)

// TrustedProxies 허용 프록시
var TrustedProxies = []string{"192.168.0.1"}

func Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("--- gin 서버가 종료되었습니다 ---")
			log.Println(err)
		}
	}()

	r := gin.Default()

	// 허용 프록시 설정
	if err := r.SetTrustedProxies(TrustedProxies); err != nil {
		panic(err)
	}

	// 미들웨어 설정
	middlewares.SetMiddleware(r)

	// 라우터 설정
	routers.SetRouters(r)

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
