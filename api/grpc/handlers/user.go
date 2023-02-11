package handlers

import (
	"context"
	"sso/api/grpc/proto"
	"sso/cmd/token"
	"sso/cmd/user"
	"sso/config/database"
	"sso/config/errors"
	"sso/config/errors/status"
)

// GetUserInformation 사용자 정보 조회
func (s *Server) GetUserInformation(c context.Context, in *proto.GetUserInformationRequest) (*proto.GetUserInformationResponse, error) {
	// 기본 정보 조회
	reqToken := in.GetToken()
	reqSessionId := in.GetSessionId()

	payload, err := token.VerifyToken(reqToken, reqSessionId)
	if err != nil {
		return nil, err
	}

	// 사용자 조회
	var userObj user.User
	userId := payload.ServerUser
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	db.Where("id = ?", userId).First(&userObj)
	if userObj.ID == 0 {
		return nil, errors.New(status.NotFoundUser)
	}
	return &proto.GetUserInformationResponse{
		User: &proto.UserInformation{
			Email: userObj.Email,
		},
	}, nil
}
