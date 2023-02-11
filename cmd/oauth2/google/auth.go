package google

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sso/config/errors"
	"sso/config/errors/status"
)

// ConfigGoogle 구글 클라이언트 설정
func ConfigGoogle() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("AUTH_GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH_GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("AUTH_GOOGLE_REDIRECT_URL"),
		Scopes:       []string{AuthGoogleCreateUserUrl},
		Endpoint:     google.Endpoint,
	}
	return conf
}

// GetEmail 구글 이메일 조회
func GetEmail(token string) (string, error) {
	var data Response

	reqURL, err := url.Parse(AuthGoogleGetUserInfoUrl)
	if err != nil {
		return "", errors.New(status.GoogleEmailRetrievalError)
	}
	resp := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{"Authorization": {fmt.Sprintf("Bearer %s", token)}},
	}
	req, err := http.DefaultClient.Do(resp)
	if err != nil {
		return "", errors.New(status.GoogleEmailRetrievalError)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", errors.New(status.GoogleEmailRetrievalError)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", errors.New(status.GoogleEmailRetrievalError)
	}
	return data.Email, nil
}
