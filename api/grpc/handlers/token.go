package handlers

import (
	"context"
	"sso/api/grpc/proto"
	"sso/cmd/client"
	"sso/cmd/token"
	"sso/cmd/user"
	"sso/config/database"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/security"
	"strconv"
)

// GenerateToken 토큰 생성
func (s *Server) GenerateToken(c context.Context, in *proto.GenerateTokenRequest) (*proto.GenerateTokenResponse, error) {
	var userObj user.User
	var groupObj user.Group
	var clientObj client.Client
	var groupUserObj user.GroupUser

	// 데이터 조회
	reqClientId := in.GetClientId()
	reqSessionId := in.GetSessionId()
	reqClientUser := in.GetClientUser()
	reqServerUser := in.GetServerUser()

	// 1. 클라이언트 존재 여부 확인
	clientIdStr, err := security.AESCipherDecrypt(reqClientId)
	if err != nil {
		return nil, err
	}
	clientId, err := strconv.Atoi(clientIdStr)
	if err != nil {
		return nil, err
	}

	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	db.Where("id = ?", clientId).First(&clientObj)
	if clientObj.ID == 0 {
		return nil, errors.New(status.NotFoundClient)
	}

	// 2. 클라이언트 시크릿키로 암호화된 user_id 복호화
	secretKey := clientObj.SecretKey
	userIdStr, err := security.AESCipherDecrypt(reqServerUser, security.CipherConfig{AESCipherKey: secretKey})
	if err != nil {
		return nil, err
	}
	userId, err := strconv.Atoi(userIdStr)

	db.Where(&user.GroupUser{UserID: userId}).First(&groupUserObj)
	db.Where("id = ?", groupUserObj.ID).First(&groupObj)
	db.Where("id = ?", userId).First(&userObj)

	// 3. 사용자 인증 여부 확인
	if !userObj.Verified {
		return nil, errors.New(status.UnauthorizedUser)
	}

	// 4. 토큰 생성
	tokenPair, err := token.GetToken(token.Payload{
		ServerUser: userId,
		SessionId:  reqSessionId,
		Group:      groupObj.UUID,
		ClientUser: reqClientUser,
	})
	return &proto.GenerateTokenResponse{
		AccessToken: &proto.Token{
			Token:     tokenPair.AccessToken.Token,
			ExpiredAt: tokenPair.AccessToken.ExpiredAt,
		},
		RefreshToken: &proto.Token{
			Token:     tokenPair.RefreshToken.Token,
			ExpiredAt: tokenPair.RefreshToken.ExpiredAt,
		},
	}, nil
}

// VerifyToken 토큰 유효성 검사
func (s *Server) VerifyToken(c context.Context, in *proto.VerifyTokenRequest) (*proto.VerifyTokenResponse, error) {
	// 기본 정보 조회
	reqToken := in.GetToken()
	reqSessionId := in.GetSessionId()

	payload, err := token.VerifyToken(reqToken, reqSessionId)
	if err != nil {
		return nil, err
	}
	return &proto.VerifyTokenResponse{
		User:  payload.ClientUser,
		Group: payload.Group,
	}, nil
}

// RefreshToken 토큰 재발급
func (s *Server) RefreshToken(c context.Context, in *proto.RefreshTokenRequest) (*proto.RefreshTokenResponse, error) {
	// 기본 정보 조회
	reqToken := in.GetToken()
	reqSessionId := in.GetSessionId()

	// 토큰 재발급
	if resp, err := token.RefreshToken(reqToken, reqSessionId); err != nil {
		return nil, err
	} else {
		return &proto.RefreshTokenResponse{
			AccessToken: &proto.Token{
				Token:     resp.AccessToken.Token,
				ExpiredAt: resp.AccessToken.ExpiredAt,
			},
			RefreshToken: &proto.Token{
				Token:     resp.RefreshToken.Token,
				ExpiredAt: resp.RefreshToken.ExpiredAt,
			},
		}, nil
	}
}
