package common

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

type Email struct {
	From     string
	Pwd      string
	To       []string
	SmtpHost string
	SmtpPort string
	Subject  string
	Content  []byte
}

func NewEmail(from, pwd string, to []string, smtpHost, smtpPort string, subject string, content []byte) (*Email, error) {
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
	auth := smtp.PlainAuth("", e.From, e.Pwd, e.SmtpHost)
	addr := fmt.Sprintf("%s:%s", e.SmtpHost, e.SmtpPort)

	conn, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("erro ao conectar: %w", err)
	}
	defer conn.Close()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         e.SmtpHost,
	}
	if err = conn.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("erro ao iniciar TLS: %w", err)
	}

	if err = conn.Auth(auth); err != nil {
		return fmt.Errorf("erro de autenticação: %w", err)
	}

	if err = conn.Mail(e.From); err != nil {
		return err
	}

	for _, addr := range e.To {
		if err = conn.Rcpt(addr); err != nil {
			return fmt.Errorf("erro ao definir destinatário %s: %w", addr, err)
		}
	}

	w, err := conn.Data()
	if err != nil {
		return fmt.Errorf("erro ao obter writer para dados: %w", err)
	}
	defer w.Close()

	headers := fmt.Sprintf("Subject: %s\r\n", e.Subject)
	headers += fmt.Sprintf("From: %s\r\n", e.From)
	headers += fmt.Sprintf("To: %s\r\n", e.To[0])
	headers += "Content-Type: text/plain; charset=utf-8\r\n"
	headers += "\r\n"

	if _, err = w.Write([]byte(headers + string(e.Content))); err != nil {
		return fmt.Errorf("erro ao escrever dados do e-mail: %w", err)
	}

	return nil
}
