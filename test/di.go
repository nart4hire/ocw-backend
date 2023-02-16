//go:build wireinject
// +build wireinject

package test

import (
	"github.com/google/wire"

	"gitlab.informatika.org/ocw/ocw-backend/handler"
	"gitlab.informatika.org/ocw/ocw-backend/middleware"
	"gitlab.informatika.org/ocw/ocw-backend/repository"
	"gitlab.informatika.org/ocw/ocw-backend/routes"
	"gitlab.informatika.org/ocw/ocw-backend/service"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/test/db"
	"gitlab.informatika.org/ocw/ocw-backend/utils"

	"gitlab.informatika.org/ocw/ocw-backend/utils/app"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

func CreateServer(logger logger.Logger, envTest *env.Environment) (app.Server, error) {
	wire.Build(
		utils.UtilSetTest,
		repository.RepositoryBasicSet,
		handler.HandlerSet,
		middleware.MiddlewareSet,
		routes.RoutesSet,
		service.ServiceTestSet,
		db.DbTestSet,
	)

	return nil, nil
}
