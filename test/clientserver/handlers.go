package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "sso/api/grpc/proto"
	"sso/pkg/http"
	"sso/pkg/security"
	"time"
)

// testMainPageHandler [테스트] 메인 페이지
func testMainPageHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

// testCreateClientHandler [테스트] 클라이언트 생성
func testCreateClientHandler(c *gin.Context) {
	var payload map[string]string

	url := ServerHost + "/v1/client/"
	bodyMap := map[string]string{
		"name":         "tester",
		"callback_uri": "http://localhost:8081/callback",
	}
	body, _ := json.Marshal(&bodyMap)
	resp, _ := http.Request("POST", url, bytes.NewReader(body))
	if resp.StatusCode != 201 {
		c.JSON(resp.StatusCode, "Exists")
		return
	}

	if err := json.Unmarshal([]byte(resp.Body), &payload); err != nil {
		c.JSON(500, nil)
		return
	}

	db := GetDatabase()
	db.Create(&Config{Key: "client_id", Value: payload["client_id"]})
	db.Create(&Config{Key: "secret_key", Value: payload["secret_key"]})
	c.JSON(201, payload)
}

// testCreateGroupHandler [테스트] 그룹 생성
func testCreateGroupHandler(c *gin.Context) {
	var clientId Config

	db := GetDatabase()
	db.Where(&Config{Key: "client_id"}).First(&clientId)

	url := ServerHost + "/v1/client/group/"
	bodyMap := map[string]interface{}{
		"name":      "default",
		"client_id": int(clientId.ID),
	}
	body, _ := json.Marshal(&bodyMap)
	resp, _ := http.Request("POST", url, bytes.NewReader(body))
	if resp.StatusCode != 201 {
		c.JSON(resp.StatusCode, "Exists")
		return
	}
	c.JSON(201, map[string]string{"group_name": "default"})
}

// testLoginHandler [테스트] 로그인
func testLoginHandler(c *gin.Context) {
	var clientId Config
	var secretKey Config

	_, err := verifyToken(c)
	if err == nil {
		c.JSON(400, map[string]string{"message": "이미 로그인된 상태입니다"})
		return
	}

	db := GetDatabase()
	db.Where(&Config{Key: "client_id"}).First(&clientId)
	db.Where(&Config{Key: "secret_key"}).First(&secretKey)
	group, _ := security.AESCipherEncrypt("default", security.CipherConfig{AESCipherKey: secretKey.Value})

	url := fmt.Sprintf(ServerHost+"?code=%s&group=%s", clientId.Value, group)
	c.JSON(200, map[string]string{"url": url})
}

// testCallbackHandler [테스트] 콜백
func testCallbackHandler(c *gin.Context) {
	var clientId Config
	var secretKey Config

	authUser := c.Query("authuser")

	db := GetDatabase()
	db.Where(&Config{Key: "client_id"}).First(&clientId)
	db.Where(&Config{Key: "secret_key"}).First(&secretKey)

	// 토큰 요청
	conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	sso := pb.NewServicesSSOClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	token, _ := sso.GenerateToken(ctx, &pb.GenerateTokenRequest{
		ClientId:   clientId.Value,
		SessionId:  "session-id",
		ServerUser: authUser,
		ClientUser: "jake",
		Group:      "default",
	})
	accessToken := token.AccessToken.Token
	refreshToken := token.RefreshToken.Token

	// 이메일 주소 요청
	user, _ := sso.GetUserInformation(ctx, &pb.GetUserInformationRequest{
		Token:     accessToken,
		SessionId: "session-id",
	})
	email := user.User.Email

	c.SetCookie("access_token", accessToken, 60, "", "", true, true)
	c.SetCookie("refresh_token", refreshToken, 60, "", "", true, true)
	c.SetCookie("session_id", "session-id", 60, "", "", true, true)

	// 리다이렉트
	redirectUrl := "/?user=" + email
	c.Redirect(302, redirectUrl)
}

// verifyToken 토큰 검증
func verifyToken(c *gin.Context) (*pb.VerifyTokenResponse, error) {
	token, _ := c.Cookie("access_token")
	sessionId, _ := c.Cookie("session_id")

	// 토큰 검증
	conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	sso := pb.NewServicesSSOClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return sso.VerifyToken(ctx, &pb.VerifyTokenRequest{
		Token:     token,
		SessionId: sessionId,
	})
}

// testStatusCheckHandler [테스트] 상태 확인
func testStatusCheckHandler(c *gin.Context) {
	token, _ := c.Cookie("access_token")

	// 토큰 검증
	conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	sso := pb.NewServicesSSOClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := verifyToken(c)
	if err == nil {
		respUser, _ := sso.GetUserInformation(ctx, &pb.GetUserInformationRequest{
			Token:     token,
			SessionId: "session-id",
		})
		email := respUser.User.Email
		c.JSON(200, map[string]string{"status": email + "로 로그인된 상태입니다"})
	} else {
		c.JSON(200, map[string]string{"status": "로그인되지 않은 상태입니다"})
	}
}

// testLogoutHandler [테스트] 로그아웃
func testLogoutHandler(c *gin.Context) {
	c.SetCookie("access_token", "", 60, "", "", true, true)
	c.SetCookie("refresh_token", "", 60, "", "", true, true)
	c.SetCookie("session_id", "", 60, "", "", true, true)
	c.JSON(204, nil)
}
