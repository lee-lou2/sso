package token

// TokenPair 토큰 결과 값
type TokenPair struct {
	AccessToken  Token `json:"access_token"`
	RefreshToken Token `json:"refresh_token"`
}

// Token 토큰
type Token struct {
	Type      string `json:"type"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}
