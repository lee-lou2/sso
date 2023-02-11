package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	netUrl "net/url"
	"os"
	"sso/pkg/cache/methods"
	"sso/pkg/datatype/convert"
	"sso/pkg/http"
)

// MessageTemplateObjectLink 메시지 템플릿 링크
type MessageTemplateObjectLink struct {
	WebUrl       string `json:"web_url"`
	MobileWebUrl string `json:"mobile_web_url"`
}

// MessageTemplateObject 메시지 템플릿
type MessageTemplateObject struct {
	ObjectType string                    `json:"object_type"`
	Text       string                    `json:"text"`
	Link       MessageTemplateObjectLink `json:"link"`
}

// ExternalAPIKaKaOClient 카카오 클라이언트 정보
type ExternalAPIKaKaOClient struct {
	ClientId    string
	Code        string
	RedirectUri string
	Scopes      string
}

// SimpleMessageTemplate 기본 메시지 템플릿
func SimpleMessageTemplate(message string) MessageTemplateObject {
	templateObjectLink := &MessageTemplateObjectLink{
		WebUrl:       "",
		MobileWebUrl: "",
	}
	templateObject := &MessageTemplateObject{
		ObjectType: "text",
		Text:       message,
		Link:       *templateObjectLink,
	}
	return *templateObject
}

// setKaKaOToken 토큰 저장
func setKaKaOToken(tokenValue map[string]interface{}) error {
	if _, hasError := tokenValue["error_code"]; hasError {
		// 오류
		err := fmt.Errorf("토큰 조회를 실패하였습니다 : %s", tokenValue["error_code"].(string))
		return err
	}
	if _, hasAccessKey := tokenValue["access_token"]; hasAccessKey {
		if _, hasRefreshKey := tokenValue["refresh_token"]; hasRefreshKey {
			// 리프레시 토큰 저장
			if err := methods.SetValue(
				"kakao_refresh_token",
				tokenValue["refresh_token"].(string),
				int(tokenValue["refresh_token_expires_in"].(float64)),
			); err != nil {
				return err
			}
		}
		// 액세스 토큰 저장
		if err := methods.SetValue(
			"kakao_access_token",
			tokenValue["access_token"].(string),
			int(tokenValue["expires_in"].(float64)),
		); err != nil {
			return err
		}
	} else {
		err := fmt.Errorf("토큰 값이 포함되어있지 않습니다")
		return err
	}
	return nil
}

// CreateKaKaOToken 토큰 생성
func CreateKaKaOToken() error {
	clientId := os.Getenv("KAKAO_CLIENT_ID")
	redirectUri := os.Getenv("KAKAO_REDIRECT_URI")
	code := os.Getenv("KAKAO_CODE")

	url := "https://kauth.kakao.com/oauth/token"

	params := netUrl.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", clientId)
	params.Add("redirect_uri", redirectUri)
	params.Add("code", code)

	resp, err := http.Request(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
		&http.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
	)
	if err != nil {
		return err
	}
	response, err := convert.StringToMap(&resp.Body)
	if err != nil {
		return err
	}
	if err := setKaKaOToken(response); err != nil {
		return err
	}
	return nil
}

// RefreshKaKaoToken 토큰 재발급
func RefreshKaKaoToken() error {
	clientId := os.Getenv("KAKAO_CLIENT_ID")
	refreshToken, _ := methods.GetValue("kakao_refresh_token")

	url := "https://kauth.kakao.com/oauth/token"

	params := netUrl.Values{}
	params.Add("grant_type", "refresh_token")
	params.Add("client_id", clientId)
	params.Add("refresh_token", refreshToken)

	resp, err := http.Request(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
	)
	if err != nil {
		return err
	}
	response, err := convert.StringToMap(&resp.Body)

	if err != nil {
		return err
	}
	if err := setKaKaOToken(response); err != nil {
		return err
	}
	return nil
}

// SendKaKaOMessageToMe 나에게 메시지 보내기
func SendKaKaOMessageToMe(message *string) error {
	templateObject := SimpleMessageTemplate(*message)
	templateObjectBytes, _ := json.Marshal(templateObject)
	templateObjectString := string(templateObjectBytes)

	url := "https://kapi.kakao.com/v2/api/talk/memo/default/send"

	accessToken, _ := methods.GetValue("kakao_access_token")

	params := netUrl.Values{}
	params.Add("template_object", templateObjectString)

	resp, err := http.Request(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
		&http.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
		&http.Header{Key: "Authorization", Value: "Bearer " + accessToken},
	)
	if err != nil {
		return err
	}
	log.Println(resp.Body)
	return nil
}

// SendKaKaOMessageToFriend 친구에게 메시지 보내기
func SendKaKaOMessageToFriend(message *string, friendUuid string) error {
	if friendUuid == "" {
		_friendUuid := os.Getenv("KAKAO_DEFAULT_FRIEND_UUID")
		friendUuid = _friendUuid
	}
	templateObject := SimpleMessageTemplate(*message)
	templateObjectBytes, _ := json.Marshal(templateObject)
	templateObjectString := string(templateObjectBytes)

	url := "https://kapi.kakao.com/v1/api/talk/friends/message/default/send"
	accessToken, _ := methods.GetValue("kakao_access_token")

	params := netUrl.Values{}
	params.Add("receiver_uuids", fmt.Sprintf(`["%s"]`, friendUuid))
	params.Add("template_object", templateObjectString)

	resp, err := http.Request(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
		&http.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
		&http.Header{Key: "Authorization", Value: "Bearer " + accessToken},
	)
	if err != nil {
		return err
	}
	log.Println(resp.Body)
	return nil
}
