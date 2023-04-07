package guard

import (
	"net/http"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type GuardBuilder struct {
	token.TokenUtil
	logger.Logger
	wrapper.WrapperUtil
}

func NewBuilder(
	token token.TokenUtil,
	logger logger.Logger,
	wrapper wrapper.WrapperUtil,
) *GuardBuilder {
	return &GuardBuilder{
		token,
		logger,
		wrapper,
	}
}

func (g *GuardBuilder) Build(role ...user.UserRole) func(http.Handler) http.Handler {
	handler := &GuardMiddleware{
		Token:       g.TokenUtil,
		Role:        role,
		Logger:      g.Logger,
		WrapperUtil: g.WrapperUtil,
	}

	return handler.Handle
}
