//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"gitlab.informatika.org/ocw/ocw-backend/app"
	"gitlab.informatika.org/ocw/ocw-backend/handler"
	"gitlab.informatika.org/ocw/ocw-backend/middleware"
	"gitlab.informatika.org/ocw/ocw-backend/provider"
	"gitlab.informatika.org/ocw/ocw-backend/repository"
	"gitlab.informatika.org/ocw/ocw-backend/routes"
	"gitlab.informatika.org/ocw/ocw-backend/service"
	"gitlab.informatika.org/ocw/ocw-backend/utils"
)

func CreateServer() (app.Server, error) {
	wire.Build(
		utils.UtilSet,
		repository.RepositorySet,
		handler.HandlerSet,
		middleware.MiddlewareSet,
		routes.RoutesSet,
		service.ServiceSet,
		provider.ProviderSet,
		app.AppSet,
	)

	return nil, nil
}
