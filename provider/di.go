package provider

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail"
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail/smtp"
	"gitlab.informatika.org/ocw/ocw-backend/provider/redis"
	"gitlab.informatika.org/ocw/ocw-backend/provider/storage"
)

var ProviderTestSet = wire.NewSet(
	// Provider
	smtp.New,
	mail.NewQueue,

	wire.Bind(new(mail.MailQueue), new(*mail.MailQueueImpl)),
	wire.Bind(new(mail.MailProvider), new(*smtp.SmtpMailProvider)),

	// Redis utility
	wire.Bind(new(redis.Redis), new(*redis.RedisImpl)),
	redis.NewRedisEnv,

	// Storage utility
	wire.Bind(new(storage.Storage), new(*storage.S3Storage)),
	storage.NewS3,
)

var ProviderSet = wire.NewSet(
	ProviderTestSet,

	// Database utility
	wire.Bind(new(db.Database), new(*db.DatabaseImpl)),
	db.NewPostgresEnv,
)
