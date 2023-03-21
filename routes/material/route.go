package material

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/material"
)

type MaterialRoutes struct {
	material.MaterialHandler
}

func (c MaterialRoutes) Register(r chi.Router) {
	r.Route("/material/{material-id}", func(r chi.Router) {
		// Add
		r.Post("/content", c.AddContent)

		// Delete
		r.Delete("/", c.DeleteMaterial)
		r.Delete("/content/{content-id}", c.DeleteContent)
	})
}
