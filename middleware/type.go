package middleware

import (
	"net/http"
)

type MiddlewareCollection interface {
	Register() []Middleware
}

type Middleware interface {
	Handle(next http.Handler) http.Handler
}
