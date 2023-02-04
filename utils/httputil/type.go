package httputil

import "net/http"

type HttpUtil interface {
	WriteSuccessJson(w http.ResponseWriter, payload interface{}) error
	WriteJson(w http.ResponseWriter, httpCode int, payload interface{}) error
	ParseJson(r *http.Request, output interface{}) error
}
