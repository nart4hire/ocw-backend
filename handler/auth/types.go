package auth

import "net/http"

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	EmailVerify(w http.ResponseWriter, r *http.Request)
}
