package guard

import (
	"net/http"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type GuardBuilder struct {
	GuardMiddleware
}

func NewBuilder(
	token token.TokenUtil,
	logger logger.Logger,
	wrapper wrapper.WrapperUtil,
) *GuardBuilder {
	return &GuardBuilder{
		GuardMiddleware{
			Token:       token,
			Role:        []user.UserRole{},
			Logger:      logger,
			WrapperUtil: wrapper,
		},
	}
}

func (g *GuardBuilder) AddRole(role ...user.UserRole) *GuardBuilder {
	g.GuardMiddleware.Role = role
	return g
}

func (g *GuardBuilder) Build() func(http.Handler) http.Handler {
	return g.GuardMiddleware.Handle
}

func (g *GuardBuilder) BuildSimple(role user.UserRole) func(http.Handler) http.Handler {
	g.AddRole(role)
	return g.Build()
}
