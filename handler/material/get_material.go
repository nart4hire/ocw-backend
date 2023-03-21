package material

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
)

func (m MaterialHandlerImpl) GetMaterial(w http.ResponseWriter, r *http.Request) {
	courseId := chi.URLParam(r, "course-id")

	if r.Header.Get("Content-Type") != "application/json" {
		payload := m.WrapperUtil.ErrorResponseWrap("this service only receive json input", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	if courseId == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("course id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	result, err := m.MaterialService.Get(courseId)

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

	responsePayload := m.WrapperUtil.SuccessResponseWrap(result)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)
}
