package app

import "net/http"

func (l *HttpServer) GetServer() http.Handler {
	return l.server
}
