package material

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	authToken "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/material"
)

func (m MaterialHandlerImpl) CreateMaterial(w http.ResponseWriter, r *http.Request) {
	courseId := chi.URLParam(r, "id")

	if courseId == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("course id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	user, ok := r.Context().Value(guard.UserContext).(authToken.UserClaim)

	if !ok {
		payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	id, err := m.MaterialService.Create(
		courseId,
		user.Email,
	)

	if err != nil {
		respErr, ok := err.(web.ResponseError)
		if ok {
			payload := m.WrapperUtil.ErrorResponseWrap(respErr.Error(), respErr)

			if respErr.Code != "NOT_OWNER" {
				m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			} else {
				m.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
			}
		} else {
			payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
			m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		}
		return
	}

	responsePayload := m.WrapperUtil.SuccessResponseWrap(material.CreateMaterialResponse{
		MaterialId: id,
	})
	m.HttpUtil.WriteSuccessJson(w, responsePayload)
}
