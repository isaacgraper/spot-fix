package common

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
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

func newEmail(from, pwd string, to []string, smtpHost string, smtpPort int, subject string, content []byte) (*Email, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[email] error loading .env file: %v", err)
	}

	password := os.Getenv("EMAIL_APP_PASSWORD")
	if password == "" {
		log.Fatal("[email] EMAIL_APP_PASSWORD must be set in .env file")
	}

	

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
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", e.From, e.To[0], e.Subject, e.Content)

	auth := smtp.PlainAuth("", e.From, e.Pwd, e.SmtpHost)

	smtpAddr := fmt.Sprintf("%s:%d", e.SmtpHost, e.SmtpPort)

	err := smtp.SendMail(smtpAddr, auth, e.From, e.To, []byte(msg))
	if err != nil {
		return fmt.Errorf("[email] error sending email: %w", err)
	}

	return nil
}
