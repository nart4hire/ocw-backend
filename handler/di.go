package handler

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/handler/admin"
	"gitlab.informatika.org/ocw/ocw-backend/handler/auth"
	"gitlab.informatika.org/ocw/ocw-backend/handler/common"
	"gitlab.informatika.org/ocw/ocw-backend/handler/course"
	"gitlab.informatika.org/ocw/ocw-backend/handler/material"
	"gitlab.informatika.org/ocw/ocw-backend/handler/reset"
	"gitlab.informatika.org/ocw/ocw-backend/handler/swagger"
)

var HandlerSet = wire.NewSet(
	// Common
	wire.Struct(new(common.CommonHandlerImpl), "*"),
	wire.Bind(new(common.CommonHandler), new(*common.CommonHandlerImpl)),

	// Swagger
	wire.Struct(new(swagger.SwaggerHandlerImpl), "*"),
	wire.Bind(new(swagger.SwaggerHandler), new(*swagger.SwaggerHandlerImpl)),

	// Admin
	wire.Struct(new(admin.AdminHandlerImpl), "*"),
	wire.Bind(new(admin.AdminHandler), new(*admin.AdminHandlerImpl)),

	// Auth
	wire.Struct(new(auth.AuthHandlerImpl), "*"),
	wire.Bind(new(auth.AuthHandler), new(*auth.AuthHandlerImpl)),

	// Reset
	wire.Struct(new(reset.ResetHandlerImpl), "*"),
	wire.Bind(new(reset.ResetHandler), new(*reset.ResetHandlerImpl)),

	// Course
	wire.Struct(new(course.CourseHandlerImpl), "*"),
	wire.Bind(new(course.CourseHandler), new(*course.CourseHandlerImpl)),

	// Material
	wire.Struct(new(material.MaterialHandlerImpl), "*"),
	wire.Bind(new(material.MaterialHandler), new(*material.MaterialHandlerImpl)),
)
