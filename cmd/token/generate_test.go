package token

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	payload := Payload{
		SessionId:  "12345",
		Group:      "admin",
		ServerUser: 1,
		ClientUser: "client1",
	}

	tokens, err := GetToken(payload)
	if err != nil {
		t.Errorf("GetToken failed with error: %v", err)
	}

	if tokens.AccessToken.Type != AccessTokenString {
		t.Errorf("Access token type is incorrect, expected: %s, got: %s", AccessTokenString, tokens.AccessToken.Type)
	}

	if tokens.RefreshToken.Type != RefreshTokenString {
		t.Errorf("Refresh token type is incorrect, expected: %s, got: %s", RefreshTokenString, tokens.RefreshToken.Type)
	}
}

func TestRefreshToken(t *testing.T) {
	payload := Payload{
		SessionId:  "12345",
		Group:      "admin",
		ServerUser: 1,
		ClientUser: "client1",
	}

	tokens, err := GetToken(payload)
	if err != nil {
		t.Errorf("GetToken failed with error: %v", err)
	}

	refreshedTokens, err := RefreshToken(tokens.RefreshToken.Token, payload.SessionId)
	if err != nil {
		t.Errorf("RefreshToken failed with error: %v", err)
	}

	if refreshedTokens.AccessToken.Type != AccessTokenString {
		t.Errorf("Access token type is incorrect, expected: %s, got: %s", AccessTokenString, refreshedTokens.AccessToken.Type)
	}

	if refreshedTokens.RefreshToken.Type != RefreshTokenString {
		t.Errorf("Refresh token type is incorrect, expected: %s, got: %s", RefreshTokenString, refreshedTokens.RefreshToken.Type)
	}
}
