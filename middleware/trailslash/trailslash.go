package trailslash

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type TrailSlashMiddleware struct{}

func (TrailSlashMiddleware) Handle(next http.Handler) http.Handler {
	return middleware.RedirectSlashes(next)
}
