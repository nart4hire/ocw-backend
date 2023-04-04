package material

import "net/http"

type MaterialHandler interface {
	AddContent(w http.ResponseWriter, r *http.Request)
	DeleteContent(w http.ResponseWriter, r *http.Request)

	CreateMaterial(w http.ResponseWriter, r *http.Request)
	DetailMaterial(w http.ResponseWriter, r *http.Request)
	DeleteMaterial(w http.ResponseWriter, r *http.Request)
	GetMaterial(w http.ResponseWriter, r *http.Request)
}
