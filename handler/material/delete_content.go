package material

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	authToken "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
)

// Index godoc
//
//	@Tags			content
//	@Summary		Delete Content
//	@Description	Delete content of material
//	@Produce		json
//	@Accept			json
//	@Param			Authorization	header		string	true	"Access token"
//	@Param			id				path		string	true	"Material id"	Format(uuid)
//	@Param			content-id		path		string	true	"Content id"	Format(uuid)
//	@Success		200				{object}	web.BaseResponse
//	@Router			/material/{id}/content/{content-id} [delete]
func (m MaterialHandlerImpl) DeleteContent(w http.ResponseWriter, r *http.Request) {
	materialIdUnparsed := chi.URLParam(r, "material-id")

	if materialIdUnparsed == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("material id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	materialId, err := uuid.Parse(materialIdUnparsed)
	if err != nil {
		payload := m.WrapperUtil.ErrorResponseWrap("material id is invalid", nil)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	contentIdUnparsed := chi.URLParam(r, "content-id")

	if contentIdUnparsed == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("content id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	contentId, err := uuid.Parse(contentIdUnparsed)
	if err != nil {
		payload := m.WrapperUtil.ErrorResponseWrap("material id is invalid", nil)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	user, ok := r.Context().Value(guard.UserContext).(authToken.UserClaim)

	if !ok {
		payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	err = m.MaterialContentService.DeleteContent(
		materialId,
		user.Email,
		contentId,
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

	responsePayload := m.WrapperUtil.SuccessResponseWrap(nil)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)
}
