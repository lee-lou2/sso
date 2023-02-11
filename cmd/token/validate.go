package token

import (
	"github.com/golang-jwt/jwt"
	"os"
	"sso/config/errors"
	"sso/config/errors/status"
)

// AdditionalValues 선택적 검증 데이터
type AdditionalValues struct {
	TokenType string
}

// VerifyToken 토큰 검증
func VerifyToken(token, sessionId string, additional ...AdditionalValues) (*Payload, error) {
	// 1. 토큰 유효성 검사
	claims, err := verifyTokenBindClaims(token)
	if err != nil {
		return nil, err
	}

	// 2. 세션 정보 일치 여부 확인
	if err := verifySession(claims, sessionId); err != nil {
		return nil, err
	}

	// 3. 캐시 존재 여부 확인
	if claims["server"] == nil {
		return nil, errors.New(status.TokenMissingRequiredInfo)
	}
	userId := int(claims["server"].(float64))
	tokenType := claims["type"].(string)
	if err := ExistsTokenInCache(userId, tokenType, sessionId); err != nil {
		return nil, err
	}

	// 4. 추가 검증
	if len(additional) == 1 {
		if additional[0].TokenType != "" && tokenType != additional[0].TokenType {
			return nil, errors.New(status.InvalidTokenType)
		}
	} else {
		if tokenType != AccessTokenString {
			return nil, errors.New(status.InvalidTokenType)
		}
	}
	return &Payload{
		ServerUser: userId,
		SessionId:  sessionId,
		Group:      claims["group"].(string),
		ClientUser: claims["client"].(string),
	}, nil
}

// verifyToken 토큰 검증 후 정보 반환
func verifyTokenBindClaims(tokenValue string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(status.TokenValidationError)
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, errors.New(status.TokenVerificationFailure)
	}
	if !token.Valid {
		return nil, errors.New(status.InvalidToken)
	}
	return token.Claims.(jwt.MapClaims), nil
}

// verifySession 토큰 내 세션 확인
func verifySession(claims map[string]interface{}, sessionId string) error {
	if claims["session_id"].(string) != sessionId {
		return errors.New(status.InvalidToken)
	}
	return nil
}
