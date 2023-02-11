package token

import "time"

const (
	Issuer               = "sso-service"                  // 생성자
	AccessTokenString    = "access_token"                 // 토큰 종류 - 엑세스 토큰
	RefreshTokenString   = "refresh_token"                // 토큰 종류 - 리프레시 토큰
	defaultDuration      = time.Second * 60 * 10          // 토큰 만료시간 - 기본
	accessTokenDuration  = time.Second * 60 * 5           // 토큰 만료시간 - 기본
	refreshTokenDuration = time.Second * 60 * 60 * 24 * 7 // 토큰 만료시간 - 리프레시 토큰
	userTokenSetLimit    = 3
)
