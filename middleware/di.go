package middleware

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/cleanpath"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/cors"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/log"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/recoverer"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/trailslash"
)

var middlewareCollectionSet = wire.NewSet(
	// Cleanpath
	wire.Struct(new(cleanpath.CleanPathMiddleware), "*"),

	// Cors
	wire.Struct(new(cors.CorsMiddleware), "*"),

	// Log
	wire.Struct(new(log.RequestLogMiddleware), "*"),

	// Recoverer
	wire.Struct(new(recoverer.RecovererMiddleware), "*"),

	// Trailslash
	wire.Struct(new(trailslash.TrailSlashMiddleware), "*"),

	guard.NewBuilder,
)

var MiddlewareSet = wire.NewSet(
	middlewareCollectionSet,

	wire.Struct(new(AppMiddlewares), "*"),
	wire.Bind(new(MiddlewareCollection), new(*AppMiddlewares)),
)
