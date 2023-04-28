package routes

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/routes/admin"
	"gitlab.informatika.org/ocw/ocw-backend/routes/auth"
	"gitlab.informatika.org/ocw/ocw-backend/routes/common"
	"gitlab.informatika.org/ocw/ocw-backend/routes/course"
	"gitlab.informatika.org/ocw/ocw-backend/routes/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/routes/material"
	"gitlab.informatika.org/ocw/ocw-backend/routes/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/routes/reset"
	"gitlab.informatika.org/ocw/ocw-backend/routes/swagger"
)

var routesCollectionSet = wire.NewSet(
	wire.Struct(new(common.CommonRoutes), "*"),
	wire.Struct(new(swagger.SwaggerRoutes), "*"),
	wire.Struct(new(auth.AuthRoutes), "*"),
	wire.Struct(new(admin.AdminRoutes), "*"),
	wire.Struct(new(reset.ResetRoutes), "*"),
	wire.Struct(new(course.CourseRoutes), "*"),
	wire.Struct(new(lesson.LessonRoutes), "*"),
	wire.Struct(new(material.MaterialRoutes), "*"),
	wire.Struct(new(quiz.QuizRoutes), "*"),
)

var RoutesSet = wire.NewSet(
	routesCollectionSet,

	wire.Struct(new(AppRouter), "*"),
	wire.Bind(new(RouteCollection), new(*AppRouter)),
)
