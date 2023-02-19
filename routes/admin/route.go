package admin

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/admin"
)

type AdminRoutes struct {
	admin.AdminHandler
}

func (adr AdminRoutes) Register(r chi.Router) {
	r.Route("/admin", func(r chi.Router) {
		r.Get("/user", adr.AdminHandler.GetAllUser)
		r.Get("/user/{id}", adr.AdminHandler.GetUserByEmail)
		r.Post("/user", adr.AdminHandler.AddUser)
		r.Patch("/user/{id}", adr.AdminHandler.UpdateUser)
		r.Delete("/user/{id}", adr.AdminHandler.DeleteUser)
	})
}
