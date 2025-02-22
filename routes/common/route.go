package common

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	commonHandler "gitlab.informatika.org/ocw/ocw-backend/handler/common"
)

type CommonRoutes struct {
	commonHandler.CommonHandler
}

func (cr CommonRoutes) Register(r chi.Router) {
	r.Get("/", cr.CommonHandler.Home)
	r.Get("/test", cr.CommonHandler.Home)
	r.Handle("/ping", http.HandlerFunc(cr.CommonHandler.Home))

	r.Handle("/*", http.HandlerFunc(cr.CommonHandler.NotFound))
}
