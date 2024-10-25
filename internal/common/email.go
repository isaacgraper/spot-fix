package common

import (
	"fmt"
	"log"
	"net/smtp"
)

type Email struct {
	From string
	Pwd  string
	To   []string

	SmtpHost string
	SmtpPort string

	Message string
	Content []byte
}

func NewEmail(from, pwd string, to []string, smtpHost, smtpPort string, message string, content []byte) (*Email, error) {
	return &Email{
		From:     from,
		Pwd:      pwd,
		To:       to,
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
		Message:  message,
		Content:  content,
	}, nil
}

func (e *Email) SendEmail() error {
	auth := smtp.PlainAuth("", e.From, e.Pwd, e.SmtpHost)
	addr := fmt.Sprintf("%s:%s", e.SmtpHost, e.SmtpPort)

	err := smtp.SendMail(addr, auth, e.From, e.To, e.Content)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	log.Println("[email] report sent to: ", e.To)
	return nil
}
