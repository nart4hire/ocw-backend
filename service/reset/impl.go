package reset

import (
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
	"gitlab.informatika.org/ocw/ocw-backend/repository/cache"
	"gitlab.informatika.org/ocw/ocw-backend/service/verification"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/password"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
)

type ResetServiceImpl struct {
	user.UserRepository
	cache.CacheRepository
	password.PasswordUtil
	*env.Environment
	token.TokenUtil
	verification.VerificationService
	logger.Logger
}
