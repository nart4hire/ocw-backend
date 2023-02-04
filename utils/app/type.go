package app

type Server interface {
	Start()
	ListRoute()
	Version()
	ListMiddleware()
}
