package gin

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"os"
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

	if os.Getenv("PROJECT_ENVIRONMENT") == "PROD" {
		runCertEngine(r)
	} else {
		runLocalEngine(r)
	}
}

// runLocalEngine 로컬 환경에서의 실행
func runLocalEngine(r *gin.Engine) {
	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}

// runCertEngine 운영 환경에서의 실행
func runCertEngine(r *gin.Engine) {
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("CERTS_HOST_WHITE_LIST")),
		Cache:      autocert.DirCache("./api/gin/certs"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))
}
