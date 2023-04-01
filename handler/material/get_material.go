package material

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
)

// Index godoc
//
//	@Tags			content
//	@Summary		Get materials
//	@Description	Get materials
//	@Produce		json
//	@Accept			json
//	@Param			id	path		string	true	"Course id"	example(IF3270)
//	@Success		200	{object}	web.BaseResponse{data=[]material.Material}
//	@Router			/course/{id}/materials [get]
func (m MaterialHandlerImpl) GetMaterial(w http.ResponseWriter, r *http.Request) {
	courseId := chi.URLParam(r, "id")

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
