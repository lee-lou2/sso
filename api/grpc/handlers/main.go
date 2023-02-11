package handlers

import (
	"sso/api/grpc/proto"
)

// Server gRPC 서버
type Server struct {
	proto.UnimplementedServicesSSOServer
}
