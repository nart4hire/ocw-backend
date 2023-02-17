package auth

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/auth"
)

type AuthRoutes struct {
	auth.AuthHandler
}

func (ar AuthRoutes) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", ar.AuthHandler.Login)
		r.Post("/refresh", ar.AuthHandler.Refresh)
		r.Post("/register", ar.AuthHandler.Register)
	})
}
