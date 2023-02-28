package reset

import "net/http"

type ResetHandler interface {
	Request(w http.ResponseWriter, r *http.Request)
	Confirm(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
	Test(w http.ResponseWriter, r *http.Request)
}
