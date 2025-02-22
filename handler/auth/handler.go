package auth

import (
	"gitlab.informatika.org/ocw/ocw-backend/service/auth"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/service/verification"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type AuthHandlerImpl struct {
	auth.AuthService
	httputil.HttpUtil
	wrapper.WrapperUtil
	logger.Logger
	verification.VerificationService
}
