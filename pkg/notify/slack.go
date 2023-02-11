package notify

import (
	"github.com/slack-go/slack"
	"os"
)

// SendSlack 슬랙 전송
func SendSlack(channel string, message string) error {
	// 토큰 조회
	accessToken := os.Getenv("SLACK_API_TOKEN")
	// APP 생성
	api := slack.New(accessToken)
	if _, _, err := api.PostMessage(
		channel,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	); err != nil {
		return err
	}
	return nil
}
