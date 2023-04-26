package material

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
)

// Index godoc
//
//	@Tags			content
//	@Summary		Get material detail
//	@Description	Get material detail
//	@Produce		json
//	@Accept			json
//	@Param			id	path		string	true	"Material id"	example(IF3270)
//	@Success		200	{object}	web.BaseResponse{data=material.Material}
//	@Router			/material/{id} [get]
func (m MaterialHandlerImpl) DetailMaterial(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "material-id")

	if idString == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("material id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	id, err := uuid.Parse(idString)

	if err != nil {
		// invalid uuid
		payload := m.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	res, err := m.MaterialService.GetById(id)

	if err != nil {
		respErr, ok := err.(web.ResponseError)
		if ok {
			payload := m.WrapperUtil.ErrorResponseWrap(respErr.Error(), respErr)

			m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		} else {
			payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
			m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		}
		return
	}

	responsePayload := m.WrapperUtil.SuccessResponseWrap(res)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)

}
