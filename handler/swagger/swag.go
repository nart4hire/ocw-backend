package swagger

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"gitlab.informatika.org/ocw/ocw-backend/docs"
)

func (SwaggerHandlerImpl) Swagger(w http.ResponseWriter, r *http.Request) {
	handler := httpSwagger.Handler()
	handler(w, r)
}

func (SwaggerHandlerImpl) SwaggerFile(w http.ResponseWriter, r *http.Request) {
	stream := docs.GetJsonSwagger()

	w.WriteHeader(http.StatusOK)
	stream.WriteTo(w)
}
