package material

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	authToken "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/material"
)

// Index godoc
//
//	@Tags					content
//	@Summary			Add Material
//	@Description	Add new material
//	@Produce			json
//	@Accept				json
//	@Param				Authorization header string true "Access token"
//	@Param				data body material.CreateMaterialRequest true "Material Request"
//	@Param				id path string true "Course id" example(IF3230)
//	@Success			200	{object}	web.BaseResponse{data=material.CreateMaterialResponse}
//	@Success			400	{object}	web.BaseResponse
//	@Success			401	{object}	web.BaseResponse
//	@Router				/course/{id}/material [post]
func (m MaterialHandlerImpl) CreateMaterial(w http.ResponseWriter, r *http.Request) {
	payload := material.CreateMaterialRequest{}
	courseId := chi.URLParam(r, "id")

	// START OF VALIDATE
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

	validate := validator.New()
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
	// END OF VALIDATE

	if courseId == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("course id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	if isExist, _ := m.CourseRepository.IsCourseExist(courseId); !isExist {
		payload := m.WrapperUtil.ErrorResponseWrap("course id not found", nil)
		m.HttpUtil.WriteJson(w, http.StatusNotFound, payload)
		return
	}

	user, ok := r.Context().Value(guard.UserContext).(authToken.UserClaim)

	if !ok {
		m.Logger.Error("Context is not found")
		payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	id, err := m.MaterialService.Create(
		courseId,
		user.Email,
		payload.Name,
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
