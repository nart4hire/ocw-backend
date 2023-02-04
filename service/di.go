package service

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/service/common"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger/hooks"
	"gitlab.informatika.org/ocw/ocw-backend/service/reporter"
)

var ServiceSet = wire.NewSet(
	// Common service
	wire.NewSet(
		wire.Struct(new(common.CommonServiceImpl), "*"),
		wire.Bind(new(common.CommonService), new(*common.CommonServiceImpl)),
	),

	// Logger service
	wire.NewSet(
		logger.New,
		hooks.NewHookCollection,
		wire.Struct(new(hooks.LogrusReporter), "*"),
		wire.Struct(new(logger.LogrusFormatter), "*"),
		wire.Bind(new(logger.Logger), new(*logger.LogrusLogger)),
	),

	// Reporter service
	wire.NewSet(
		reporter.New,
		wire.Bind(new(reporter.Reporter), new(*reporter.LogtailReporter)),
	),
)
