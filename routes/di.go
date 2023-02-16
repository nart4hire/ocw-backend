package routes

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/routes/auth"
	"gitlab.informatika.org/ocw/ocw-backend/routes/common"
	"gitlab.informatika.org/ocw/ocw-backend/routes/swagger"
)

var routesCollectionSet = wire.NewSet(
	wire.Struct(new(common.CommonRoutes), "*"),
	wire.Struct(new(swagger.SwaggerRoutes), "*"),
	wire.Struct(new(auth.AuthRoutes), "*"),
)

var RoutesSet = wire.NewSet(
	routesCollectionSet,

	wire.Struct(new(AppRouter), "*"),
	wire.Bind(new(RouteCollection), new(*AppRouter)),
)
