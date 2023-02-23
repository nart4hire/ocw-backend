package provider

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail"
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail/smtp"
)

var ProviderTestSet = wire.NewSet(
	// Provider
	smtp.New,
	mail.NewQueue,

	wire.Bind(new(mail.MailQueue), new(*mail.MailQueueImpl)),
	wire.Bind(new(mail.MailProvider), new(*smtp.SmtpMailProvider)),
)

var ProviderSet = wire.NewSet(
	ProviderTestSet,

	// Database utility
	wire.Bind(new(db.Database), new(*db.DatabaseImpl)),
	db.NewPostgresEnv,
)
