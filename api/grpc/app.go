package grpc

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"sso/api/grpc/handlers"
	pb "sso/api/grpc/proto"
)

// Run gRPC 서버 실행
func Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("--- gRPC 서버가 종료되었습니다 ---")
			log.Println(err)
		}
	}()

	logrus.ErrorKey = "grpc.error"
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	listen, _ := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	rpc := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		)),
	)
	pb.RegisterServicesSSOServer(rpc, &handlers.Server{})
	log.Printf("gRPC 서버 실행 되었습니다")
	if err := rpc.Serve(listen); err != nil {
		log.Fatalf("gRPC 연결 실패, 오류 내용 : %v", err)
	}
}
