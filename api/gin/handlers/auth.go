package handlers

import (
	"github.com/gin-gonic/gin"
	"sso/cmd/oauth2"
	"sso/cmd/oauth2/google"
	"sso/config/errors"
	"sso/config/errors/status"
)

const (
	OAuth2ProviderParamName = "provider"
)

// OAuth2RedirectHandler 공급사 로그인 페이지로 이동
func OAuth2RedirectHandler(c *gin.Context) {
	provider := c.Param(OAuth2ProviderParamName)

	if provider == oauth2.ProviderGoogleString {
		google.OAuth2Google(c)
		return
	}
	errors.GinErrorJSON(c, errors.New(status.NotFoundProvider))
}

// OAuth2CallbackHandler 공급사 회원 가입 또는 로그인
func OAuth2CallbackHandler(c *gin.Context) {
	provider := c.Param(OAuth2ProviderParamName)

	if provider == oauth2.ProviderGoogleString {
		google.OAuth2GoogleCallback(c)
		return
	}
	errors.GinErrorTemplate(c, errors.New(status.NotFoundProvider))
}
