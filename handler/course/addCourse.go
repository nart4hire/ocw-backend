package course

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course"
)

// Index godoc
//
//	@Summary		Add a new course
//	@Description	Add a new course with the given details
//	@Tags			course
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string							true	"AddCourseToken"
//	@Param			data			body		course.AddCourseRequestPayload	true	"Add Course payload"
//	@Success		200				{object}	web.BaseResponse				"Success"
//	@Failure		400				{object}	web.BaseResponse				"Bad Request"
//	@Failure		401				{object}	web.BaseResponse				"Unauthorized"
//	@Failure		403				{object}	web.BaseResponse				"Forbidden"
//	@Failure		422				{object}	web.BaseResponse				"Unprocessable Entity"
//	@Failure		500				{object}	web.BaseResponse				"Internal Server Error"
//	@Router			/course [put]
func (c CourseHandlerImpl) AddCourse(w http.ResponseWriter, r *http.Request) {
	payload := course.AddCourseRequestPayload{}

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

	validate := validator.New()
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

	payload.AddCourseToken = token[1]
	err := c.CourseService.AddCourse(payload)

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
