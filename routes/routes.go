package routes

import (
	"gitlab.informatika.org/ocw/ocw-backend/routes/admin"
	"gitlab.informatika.org/ocw/ocw-backend/routes/auth"
	"gitlab.informatika.org/ocw/ocw-backend/routes/common"
	"gitlab.informatika.org/ocw/ocw-backend/routes/swagger"

	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
)

type AppRouter struct {
	// Routes
	swagger.SwaggerRoutes
	admin.AdminRoutes
	common.CommonRoutes
	auth.AuthRoutes

	// Utility
	Logger logger.Logger
}
