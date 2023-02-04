package httputil

import (
	"encoding/json"
	"net/http"
)

func (HttpUtilImpl) WriteJson(w http.ResponseWriter, httpCode int, payload interface{}) error {
	encoder := json.NewEncoder(w)
	w.WriteHeader(httpCode)
	return encoder.Encode(payload)
}

func (h HttpUtilImpl) WriteSuccessJson(w http.ResponseWriter, payload interface{}) error {
	return h.WriteJson(w, http.StatusOK, payload)
}
