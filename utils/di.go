package utils

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/utils/base64"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/log"
	"gitlab.informatika.org/ocw/ocw-backend/utils/password"
	"gitlab.informatika.org/ocw/ocw-backend/utils/res"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
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

	// Base64 Utility
	wire.Struct(new(base64.Base64UtilImpl), "*"),
	wire.Bind(new(base64.Base64Util), new(*base64.Base64UtilImpl)),

	// Password utility
	wire.Struct(new(password.PasswordUtilImpl), "*"),
	wire.Bind(new(password.PasswordUtil), new(*password.PasswordUtilImpl)),

	// Token utility
	wire.Struct(new(token.TokenUtilImpl), "*"),
	wire.Bind(new(token.TokenUtil), new(*token.TokenUtilImpl)),
)

var UtilSet = wire.NewSet(
	// env
	env.New,

	UtilSetTest,
)
