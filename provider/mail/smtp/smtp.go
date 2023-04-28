package smtp

import (
	"fmt"
	"net/smtp"

	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

type SmtpMailProvider struct {
	*env.Environment
	smtp.Auth
}

func New(env *env.Environment) *SmtpMailProvider {
	var auth smtp.Auth
	if env.SmtpAuthType == "CRAM" {
		auth = smtp.CRAMMD5Auth(
			env.SmtpUsername,
			env.SmtpPassword,
		)
	} else if env.SmtpAuthType == "plain" {
		auth = smtp.PlainAuth(
			env.SmtpIdentity,
			env.SmtpUsername,
			env.SmtpPassword,
			env.SmtpServer,
		)
	}

	return &SmtpMailProvider{
		Environment: env,
		Auth:        auth,
	}
}

func (s SmtpMailProvider) Send(to []string, subject string, message string) error {
	payload := fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: %s\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n"+
			"\r\n%s\r\n",
		to, subject, message,
	)

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", s.SmtpServer, s.SmtpPort),
		s.Auth,
		s.SmtpUsername,
		to,
		[]byte(payload),
	)
}
