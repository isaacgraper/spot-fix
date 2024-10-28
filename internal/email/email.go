package common

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
	// smtp provider is not working
	return nil
}
