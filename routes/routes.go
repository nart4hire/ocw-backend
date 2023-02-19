package routes

import (
	"gitlab.informatika.org/ocw/ocw-backend/routes/auth"
	"gitlab.informatika.org/ocw/ocw-backend/routes/common"
	"gitlab.informatika.org/ocw/ocw-backend/routes/reset"
	"gitlab.informatika.org/ocw/ocw-backend/routes/swagger"

	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
)

type AppRouter struct {
	// Routes
	swagger.SwaggerRoutes
	common.CommonRoutes
	auth.AuthRoutes
	reset.ResetRoutes

	// Utility
	Logger logger.Logger
}
