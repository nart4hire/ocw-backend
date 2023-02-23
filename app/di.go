package app

import "github.com/google/wire"

var AppSet = wire.NewSet(
	// app
	New,
	wire.Bind(new(Server), new(*HttpServer)),
)
