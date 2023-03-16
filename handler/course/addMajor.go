package course

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course/major/add"
)

func (c CourseHandlerImpl) AddMajor(w http.ResponseWriter, r *http.Request) {
	payload := add.AddMajorRequestPayload{}
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

	payload.AddMajorToken = token[1]
	err := c.CourseService.AddMajor(payload)

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