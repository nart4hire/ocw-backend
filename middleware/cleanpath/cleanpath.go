package cleanpath

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type CleanPathMiddleware struct{}

func (CleanPathMiddleware) Handle(next http.Handler) http.Handler {
	return middleware.CleanPath(next)
}
