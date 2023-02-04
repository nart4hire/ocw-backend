package utils

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/utils/app"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
	"gitlab.informatika.org/ocw/ocw-backend/utils/res"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

var UtilSetTest = wire.NewSet(
	// httputil utility
	wire.Struct(new(httputil.HttpUtilImpl), "*"),
	wire.Bind(new(httputil.HttpUtil), new(*httputil.HttpUtilImpl)),

	// log utility
	wire.Struct(new(log.LogUtilsImpl), "*"),
	wire.Bind(new(log.LogUtils), new(*log.LogUtilsImpl)),

	// res utility
	wire.Struct(new(res.EmbedResources), "*"),
	wire.Bind(new(res.Resource), new(*res.EmbedResources)),

	// wrapper utility
	wire.Struct(new(wrapper.WrapperUtilImpl), "*"),
	wire.Bind(new(wrapper.WrapperUtil), new(*wrapper.WrapperUtilImpl)),

	// app
	app.New,
	wire.Bind(new(app.Server), new(*app.HttpServer)),
)

var UtilSet = wire.NewSet(
	UtilSetTest,

	// env
	env.New,
)
