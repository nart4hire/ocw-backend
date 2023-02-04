//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"gitlab.informatika.org/ocw/ocw-backend/handler"
	"gitlab.informatika.org/ocw/ocw-backend/middleware"
	"gitlab.informatika.org/ocw/ocw-backend/routes"
	"gitlab.informatika.org/ocw/ocw-backend/service"
	"gitlab.informatika.org/ocw/ocw-backend/utils"

	"gitlab.informatika.org/ocw/ocw-backend/utils/app"
)

func CreateServer() (app.Server, error) {
	wire.Build(
		utils.UtilSet,
		handler.HandlerSet,
		middleware.MiddlewareSet,
		routes.RoutesSet,
		service.ServiceSet,
	)

	return nil, nil
}
