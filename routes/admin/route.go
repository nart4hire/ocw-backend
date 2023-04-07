package admin

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/admin"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type AdminRoutes struct {
	admin.AdminHandler
	*guard.GuardBuilder
}

func (adr AdminRoutes) Register(r chi.Router) {
	r.Route("/admin", func(r chi.Router) {
		r.Use(adr.GuardBuilder.Build(user.Admin))

		r.Get("/user", adr.AdminHandler.GetAllUser)
		r.Get("/user/{email}", adr.AdminHandler.GetUserByEmail)
		r.Post("/user", adr.AdminHandler.AddUser)
		r.Patch("/user/{email}", adr.AdminHandler.UpdateUser)
		r.Delete("/user/{email}", adr.AdminHandler.DeleteUser)
	})
}
