package cors

import (
	"net/http"

	"github.com/go-chi/cors"
)

type CorsMiddleware struct{}

var corsHandler = cors.Handler(cors.Options{
	AllowedOrigins:   []string{"http://*", "https://*"},
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
	AllowedHeaders:   []string{"*"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: false,
	MaxAge:           300,
})

func (CorsMiddleware) Handle(next http.Handler) http.Handler {
	return corsHandler(next)
}
