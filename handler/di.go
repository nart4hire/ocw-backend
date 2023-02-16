package handler

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/handler/auth"
	"gitlab.informatika.org/ocw/ocw-backend/handler/common"
	"gitlab.informatika.org/ocw/ocw-backend/handler/swagger"
)

var HandlerSet = wire.NewSet(
	// Common
	wire.Struct(new(common.CommonHandlerImpl), "*"),
	wire.Bind(new(common.CommonHandler), new(*common.CommonHandlerImpl)),

	// Swagger
	wire.Struct(new(swagger.SwaggerHandlerImpl), "*"),
	wire.Bind(new(swagger.SwaggerHandler), new(*swagger.SwaggerHandlerImpl)),

	// Auth
	wire.Struct(new(auth.AuthHandlerImpl), "*"),
	wire.Bind(new(auth.AuthHandler), new(*auth.AuthHandlerImpl)),
)
