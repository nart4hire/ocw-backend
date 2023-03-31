package course

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course/faculty"
)

// Index godoc
//
//	@Summary		Update a faculty 
//	@Description	Update a faculty with the given ID
//	@Tags			course
//	@Param			id				path		string								true	"Faculty ID (UUID)"
//	@Param			payload			body		faculty.UpdateFacultyRequestPayload	true	"Update Faculty Payload"
//	@Param			Authorization	header		string								true	"UpdateFacultyToken"
//	@Success		200				{object}	web.BaseResponse					"Success"
//	@Failure		400				{object}	web.BaseResponse					"Bad Request"
//	@Failure		401				{object}	web.BaseResponse					"Unauthorized"
//	@Failure		403				{object}	web.BaseResponse					"Forbidden"
//	@Failure		422				{object}	web.BaseResponse					"Unprocessable Entity"
//	@Failure		500				{object}	web.BaseResponse					"Internal Server Error"
//	@Router			/course/faculty/{id} [patch]
func (c CourseHandlerImpl) UpdateFaculty(w http.ResponseWriter, r *http.Request) {
	payload := faculty.UpdateFacultyRequestPayload{}
	validate := validator.New()

	// Validate payload
	if r.Header.Get("Content-Type") != "application/json" {
		payload := c.WrapperUtil.ErrorResponseWrap("this service only receive json input", nil)
		c.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	if err := c.HttpUtil.ParseJson(r, &payload); err != nil {
		payload := c.WrapperUtil.ErrorResponseWrap("invalid json input", err.Error())
		c.HttpUtil.WriteJson(w, http.StatusUnprocessableEntity, payload)
		return
	}

	if err := validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			payload := c.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
			c.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			return
		}

		errPayload := web.NewResponseErrorFromValidator(err.(validator.ValidationErrors))
		payload := c.WrapperUtil.ErrorResponseWrap(errPayload.Error(), errPayload)
		c.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	// Confirm Valid Website Token
	validateTokenHeader := r.Header.Get("Authorization")

	if validateTokenHeader == "" {
		payload := c.WrapperUtil.ErrorResponseWrap("token is required", nil)
		c.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
		return
	}

	token := strings.Split(validateTokenHeader, " ")

	if len(token) != 2 {
		payload := c.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		c.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	if token[0] != "Bearer" {
		payload := c.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		c.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	id, err := uuid.Parse(chi.URLParam(r, "id"))

	if err != nil {
		payload := c.WrapperUtil.ErrorResponseWrap("invalid id", nil)
		c.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	payload.UpdateFacultyToken = token[1]
	payload.ID = id
	err = c.CourseService.UpdateFaculty(payload)

	if err != nil {
		if errData, ok := err.(web.ResponseError); ok {
			payload := c.WrapperUtil.ErrorResponseWrap(errData.Error(), errData)
			c.HttpUtil.WriteJson(w, http.StatusUnauthorized, payload)
			return
		}

		c.Logger.Error(
			fmt.Sprintf("[RESET] some error happened when validating URL: %s", err.Error()),
		)
		payload := c.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		c.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	responsePayload := c.WrapperUtil.SuccessResponseWrap(nil)
	c.HttpUtil.WriteSuccessJson(w, responsePayload)
}
