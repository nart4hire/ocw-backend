package service

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/service/admin"
	"gitlab.informatika.org/ocw/ocw-backend/service/auth"
	"gitlab.informatika.org/ocw/ocw-backend/service/common"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger/hooks"
	"gitlab.informatika.org/ocw/ocw-backend/service/material"
	"gitlab.informatika.org/ocw/ocw-backend/service/reporter"
	"gitlab.informatika.org/ocw/ocw-backend/service/reset"
	"gitlab.informatika.org/ocw/ocw-backend/service/verification"
)

var ServiceTestSet = wire.NewSet(
	// Common service
	wire.NewSet(
		wire.Struct(new(common.CommonServiceImpl), "*"),
		wire.Bind(new(common.CommonService), new(*common.CommonServiceImpl)),
	),

	// Reporter service
	wire.NewSet(
		reporter.New,
		wire.Bind(new(reporter.Reporter), new(*reporter.LogtailReporter)),
	),

	// auth service
	wire.NewSet(
		wire.Struct(new(auth.AuthServiceImpl), "*"),
		wire.Bind(new(auth.AuthService), new(*auth.AuthServiceImpl)),
	),

	// admin service
	wire.NewSet(
		wire.Struct(new(admin.AdminServiceImpl), "*"),
		wire.Bind(new(admin.AdminService), new(*admin.AdminServiceImpl)),
	),

	// reset service
	wire.NewSet(
		wire.Struct(new(reset.ResetServiceImpl), "*"),
		wire.Bind(new(reset.ResetService), new(*reset.ResetServiceImpl)),
	),

	// verification service
	wire.NewSet(
		wire.Struct(new(verification.VerificationServiceImpl), "*"),
		wire.Bind(new(verification.VerificationService), new(*verification.VerificationServiceImpl)),
	),

	wire.NewSet(
		wire.Struct(new(material.MaterialContentServiceImpl), "*"),
		wire.Struct(new(material.MaterialServiceImpl), "*"),
		wire.Bind(new(material.MaterialContentService), new(*material.MaterialContentServiceImpl)),
		wire.Bind(new(material.MaterialService), new(*material.MaterialServiceImpl)),
	),
)

var ServiceSet = wire.NewSet(
	ServiceTestSet,

	// Logger service
	wire.NewSet(
		logger.New,
		hooks.NewHookCollection,
		wire.Struct(new(hooks.LogrusReporter), "*"),
		wire.Struct(new(logger.LogrusFormatter), "*"),
		wire.Bind(new(logger.Logger), new(*logger.LogrusLogger)),
	),
)
