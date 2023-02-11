FROM golang:1.19.4

# 기본 패키지 업데이트
RUN apt update -y
RUN apt upgrade -y
RUN apt-get update -y
RUN apt-get upgrade -y

# 컴파일러 설치
RUN apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Python 설치
RUN apt-get install -y python3 python3-pip
RUN pip3 install grpcio
RUN pip3 install grpcio-tools
