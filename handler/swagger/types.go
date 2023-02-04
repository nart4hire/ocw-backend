package swagger

import "net/http"

type SwaggerHandler interface {
	Swagger(w http.ResponseWriter, r *http.Request)
	SwaggerFile(w http.ResponseWriter, r *http.Request)
}
