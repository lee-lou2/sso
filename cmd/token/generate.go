package token

import (
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

// Payload 토큰 정보
type Payload struct {
	SessionId  string
	Group      string
	ServerUser int
	ClientUser string
}

// GetToken 토큰 생성
func GetToken(payload Payload) (*TokenPair, error) {
	accessToken, err := generateToken(payload, AccessTokenString)
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateToken(payload, RefreshTokenString)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}

// generateToken JWT Token 생성
func generateToken(payload Payload, tokenType string) (*Token, error) {
	duration := defaultDuration
	// 토큰 종류별 만료 시간 정의
	if tokenType == AccessTokenString {
		duration = accessTokenDuration
	} else if tokenType == RefreshTokenString {
		duration = refreshTokenDuration
	}

	respToken := jwt.New(jwt.SigningMethodHS256)

	claims := respToken.Claims.(jwt.MapClaims)

	// [필수] 사용자 정보 - 재발급에 필요한 정보
	claims["server"] = payload.ServerUser
	claims["session_id"] = payload.SessionId

	// [선택] 사용자 정보 - 재발급에 필요하지 않은 정보
	claims["client"] = payload.ClientUser
	claims["group"] = payload.Group

	// [기본] 토큰 기본 정보
	claims["iss"] = Issuer
	claims["type"] = tokenType
	expiredAt := time.Now().Add(duration).Unix()
	claims["exp"] = expiredAt

	tokenValue, err := respToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	// 캐시에 토큰 저장
	if err := SetTokenInCache(payload.ServerUser, tokenType, payload.SessionId); err != nil {
		log.Println("캐시에 토큰 저장간 오류 발생, 오류 내용 : " + err.Error())
	}

	// 토큰 반환
	return &Token{
		Type:      tokenType,
		Token:     tokenValue,
		ExpiredAt: expiredAt,
	}, nil
}

// RefreshToken 토큰 재발급
func RefreshToken(token, sessionId string) (*TokenPair, error) {
	payload, err := VerifyToken(token, sessionId, AdditionalValues{TokenType: RefreshTokenString})
	if err != nil {
		return nil, err
	}

	// 토큰 발급
	tokens, err := GetToken(*payload)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
