package notify

import (
	"fmt"
	"github.com/kevinburke/twilio-go"
	"os"
)

// SendSMSTwilio 문자 발송
func SendSMSTwilio(toNumber string, message string) error {
	fromNumber := os.Getenv("TWILIO_API_FROM_NUMBER")
	sid := os.Getenv("TWILIO_API_SID")
	token := os.Getenv("TWILIO_API_AUTH_TOKEN")

	client := twilio.NewClient(sid, token, nil)

	if _, err := client.Messages.SendMessage(
		fromNumber,
		fmt.Sprintf("+82%s", toNumber[1:]),
		message,
		nil,
	); err != nil {
		return err
	}
	return nil
}
