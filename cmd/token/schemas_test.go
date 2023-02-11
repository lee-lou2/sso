package token

import (
	"testing"
	"time"
)

const (
	TokenExpiration        = 60 * 5
	RefreshTokenExpiration = 60 * 60 * 12
)

func TestToken(t *testing.T) {
	// Test the creation of a new Token
	now := time.Now().Unix()
	token := Token{
		Type:      AccessTokenString,
		Token:     "some_random_token_string",
		ExpiredAt: now + int64(TokenExpiration),
	}

	if token.Type != AccessTokenString {
		t.Errorf("expected token type %s, but got %s", AccessTokenString, token.Type)
	}
	if token.Token != "some_random_token_string" {
		t.Errorf("expected token value %s, but got %s", "some_random_token_string", token.Token)
	}
	if token.ExpiredAt != now+int64(TokenExpiration) {
		t.Errorf("expected expired at %d, but got %d", now+int64(TokenExpiration), token.ExpiredAt)
	}
}

func TestTokenPair(t *testing.T) {
	// Test the creation of a new TokenPair
	now := time.Now().Unix()
	tokenPair := TokenPair{
		AccessToken: Token{
			Type:      AccessTokenString,
			Token:     "some_random_access_token_string",
			ExpiredAt: now + int64(TokenExpiration),
		},
		RefreshToken: Token{
			Type:      RefreshTokenString,
			Token:     "some_random_refresh_token_string",
			ExpiredAt: now + int64(RefreshTokenExpiration),
		},
	}

	if tokenPair.AccessToken.Type != AccessTokenString {
		t.Errorf("expected access token type %s, but got %s", AccessTokenString, tokenPair.AccessToken.Type)
	}
	if tokenPair.AccessToken.Token != "some_random_access_token_string" {
		t.Errorf("expected access token value %s, but got %s", "some_random_access_token_string", tokenPair.AccessToken.Token)
	}
	if tokenPair.AccessToken.ExpiredAt != now+int64(TokenExpiration) {
		t.Errorf("expected access token expired at %d, but got %d", now+int64(TokenExpiration), tokenPair.AccessToken.ExpiredAt)
	}
	if tokenPair.RefreshToken.Type != RefreshTokenString {
		t.Errorf("expected refresh token type %s, but got %s", RefreshTokenString, tokenPair.RefreshToken.Type)
	}
	if tokenPair.RefreshToken.Token != "some_random_refresh_token_string" {
		t.Errorf("expected refresh token value %s, but got %s", "some_random_refresh_token_string", tokenPair.RefreshToken.Token)
	}
}
