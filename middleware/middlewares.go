package middleware

import (
	"gitlab.informatika.org/ocw/ocw-backend/middleware/cleanpath"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/cors"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/log"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/recoverer"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/trailslash"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
)

type AppMiddlewares struct {
	// Registered Middleware
	recoverer.RecovererMiddleware
	cors.CorsMiddleware
	log.RequestLogMiddleware
	trailslash.TrailSlashMiddleware
	cleanpath.CleanPathMiddleware

	// Utility
	Logger logger.Logger
}
