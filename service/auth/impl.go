package auth

import (
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/service/verification"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/password"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
)

type AuthServiceImpl struct {
	user.UserRepository
	password.PasswordUtil
	*env.Environment
	token.TokenUtil
	verification.VerificationService
	logger.Logger
}
