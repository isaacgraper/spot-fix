package common

import (
	"fmt"
	"log"

	"gopkg.in/mail.v2"
)

type Email struct {
	From     string
	Pwd      string
	To       []string
	SmtpHost string
	SmtpPort int
	Subject  string
	Content  []byte
}

func NewEmail(from, pwd string, to []string, smtpHost string, smtpPort int, subject string, content []byte) (*Email, error) {
	return &Email{
		From:     from,
		Pwd:      pwd,
		To:       to,
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
		Subject:  subject,
		Content:  content,
	}, nil
}

func (e *Email) SendEmail() error {
	msg := mail.NewMessage()

	msg.SetHeader("From", e.From)
	msg.SetHeader("To", e.To...)
	msg.SetHeader("Subject", e.Subject)

	msg.SetBody("text/plain", string(e.Content))

	dialer := mail.NewDialer(e.SmtpHost, e.SmtpPort, e.From, e.Pwd)

	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	} else {
		log.Println("[email] email sent sucessfully")
	}
	return nil
}
