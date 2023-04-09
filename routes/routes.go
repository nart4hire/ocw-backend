package routes

import (
	"gitlab.informatika.org/ocw/ocw-backend/routes/admin"
	"gitlab.informatika.org/ocw/ocw-backend/routes/auth"
	"gitlab.informatika.org/ocw/ocw-backend/routes/common"
	"gitlab.informatika.org/ocw/ocw-backend/routes/course"
	"gitlab.informatika.org/ocw/ocw-backend/routes/material"
	"gitlab.informatika.org/ocw/ocw-backend/routes/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/routes/reset"
	"gitlab.informatika.org/ocw/ocw-backend/routes/swagger"

	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
)

type AppRouter struct {
	// Routes
	swagger.SwaggerRoutes
	admin.AdminRoutes
	common.CommonRoutes
	auth.AuthRoutes
	reset.ResetRoutes
	quiz.QuizRoutes
	course.CourseRoutes
	material.MaterialRoutes

	// Utility
	Logger logger.Logger
}
