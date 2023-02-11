package notify

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

// SendSMTPEmail 이메일 전송
func SendSMTPEmail(to string, subject string, message string) error {
	// 기본 변수
	fromEmail := os.Getenv("EMAIL_HOST_USER")
	emailUser := os.Getenv("EMAIL_HOST_USER")
	emailPassword := os.Getenv("EMAIL_HOST_PASSWORD")
	emailSMTPHost := os.Getenv("EMAIL_HOST")
	emailSMTPPortString := os.Getenv("EMAIL_PORT")
	emailSMTPPort, _ := strconv.Atoi(emailSMTPPortString)

	// 이메일 발송
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	d := gomail.NewDialer(emailSMTPHost, emailSMTPPort, emailUser, emailPassword)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
