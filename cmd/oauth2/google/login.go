package google

import (
	"github.com/gin-gonic/gin"
)

// OAuth2Google 구글 Url 조회
func OAuth2Google(c *gin.Context) {
	path := ConfigGoogle()
	url := path.AuthCodeURL(c.Query("state"))
	c.Redirect(302, url)
}
