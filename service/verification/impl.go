package verification

import (
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail"
	"gitlab.informatika.org/ocw/ocw-backend/repository/cache"
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/template"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
)

type VerificationServiceImpl struct {
	mail.MailQueue
	user.UserRepository
	*env.Environment
	template.TemplateWritterBuilder
	token.TokenUtil
	cache.CacheRepository
}
