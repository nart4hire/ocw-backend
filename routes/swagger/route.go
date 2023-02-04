package swagger

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/swagger"
)

type SwaggerRoutes struct {
	swagger.SwaggerHandler
}

func (sr SwaggerRoutes) Register(r chi.Router) {
	r.Get("/docs", sr.SwaggerHandler.SwaggerFile)
	r.Get("/docs/*", sr.SwaggerHandler.Swagger)
}
