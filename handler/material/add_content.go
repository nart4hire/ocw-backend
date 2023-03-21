package material

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	materialDomain "gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	authToken "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/material"
)

func (m MaterialHandlerImpl) AddContent(w http.ResponseWriter, r *http.Request) {
	payload := material.NewContentRequest{}
	user, ok := r.Context().Value(guard.UserContext).(authToken.UserClaim)

	materialId := chi.URLParam(r, "material-id")
	if materialId == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("material id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
	}

	materialIdUUID, err := uuid.Parse(materialId)
	if err != nil {
		payload := m.WrapperUtil.ErrorResponseWrap("material id is invalid", err)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
	}

	payload.MaterialId = materialIdUUID

	if !ok {
		payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	validate := validator.New()

	if r.Header.Get("Content-Type") != "application/json" {
		payload := m.WrapperUtil.ErrorResponseWrap("this service only receive json input", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	if err := m.HttpUtil.ParseJson(r, &payload); err != nil {
		payload := m.WrapperUtil.ErrorResponseWrap("invalid json input", err.Error())
		m.HttpUtil.WriteJson(w, http.StatusUnprocessableEntity, payload)
		return
	}

	if err := validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			payload := m.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
			m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			return
		}

		errPayload := web.NewResponseErrorFromValidator(err.(validator.ValidationErrors))
		payload := m.WrapperUtil.ErrorResponseWrap(errPayload.Error(), errPayload)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	uploadLink, err := m.MaterialContentService.AddContent(payload.MaterialId, user.Email, materialDomain.Content{
		Type: payload.Type,
		Link: payload.Link,
	})

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

	responsePayload := m.WrapperUtil.SuccessResponseWrap(material.NewContentResponse{
		UploadLink: uploadLink,
	})
	m.HttpUtil.WriteSuccessJson(w, responsePayload)
}
