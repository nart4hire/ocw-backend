package reset

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/reset"
)

type ResetRoutes struct {
	reset.ResetHandler
}

func (rr ResetRoutes) Register(r chi.Router) {
	r.Route("/reset", func(r chi.Router) {
		r.Post("/request", rr.ResetHandler.Request)
		r.Post("/confirm", rr.ResetHandler.Confirm)
		r.Post("/validate", rr.ResetHandler.Validate)
		r.Get("/test", rr.ResetHandler.Test)
	})
}
