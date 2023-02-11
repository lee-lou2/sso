package routers

import (
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	r.LoadHTMLGlob("./web/views/*")
	r.Static("/static", "./web/static")

	// HealthCheck
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Web
	setRoutersWeb(r)

	// V1
	setRoutersV1(r)
}
