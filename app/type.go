package app

import "net/http"

type Server interface {
	Start()
	ListRoute()
	Version()
	ListMiddleware()
	GetServer() http.Handler
}
