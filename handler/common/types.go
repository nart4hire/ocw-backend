package common

import "net/http"

type CommonHandler interface {
	Home(w http.ResponseWriter, r *http.Request)
}
