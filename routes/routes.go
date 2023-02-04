package routes

import (
	"gitlab.informatika.org/ocw/ocw-backend/routes/common"
	"gitlab.informatika.org/ocw/ocw-backend/routes/swagger"

	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
)

type AppRouter struct {
	common.CommonRoutes
	swagger.SwaggerRoutes

	Logger logger.Logger
}
