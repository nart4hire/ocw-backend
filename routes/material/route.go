package material

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/material"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type MaterialRoutes struct {
	material.MaterialHandler
	*guard.GuardBuilder
}

func (c MaterialRoutes) Register(r chi.Router) {
	r.Route("/material/{material-id}", func(r chi.Router) {
		r.Get("/", c.DetailMaterial)

		r.Route("/", func(r chi.Router) {
			r.Use(c.GuardBuilder.BuildSimple(user.Contributor))

			// Add
			r.Post("/content", c.AddContent)

			// Delete
			r.Delete("/", c.DeleteMaterial)
			r.Delete("/content/{content-id}", c.DeleteContent)
		})
	})
}
