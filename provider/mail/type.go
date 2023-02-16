package mail

import "context"

type Mail struct {
	To      []string
	Subject string
	Message string
}

type MailProvider interface {
	Send(to []string, subject string, message string) error
}

type MailQueue interface {
	Send(mail Mail)
	Flush()
	Start(ctx context.Context)
}
