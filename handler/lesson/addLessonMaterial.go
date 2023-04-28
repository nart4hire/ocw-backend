package lesson

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson/materials"
)

// Index godoc
//
//	@Summary		Add a new lesson material
//	@Description	Add a new lesson material with the given details
//	@Tags			lesson
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string							true	"AddLessonMaterialsToken"
//	@Param			data			body		materials.AddLessonMaterialsRequestPayload	true	"Add Lesson Material payload"
//	@Success		200				{object}	web.BaseResponse				"Success"
//	@Failure		400				{object}	web.BaseResponse				"Bad Request"
//	@Failure		401				{object}	web.BaseResponse				"Unauthorized"
//	@Failure		403				{object}	web.BaseResponse				"Forbidden"
//	@Failure		422				{object}	web.BaseResponse				"Unprocessable Entity"
//	@Failure		500				{object}	web.BaseResponse				"Internal Server Error"
//	@Router			/lesson/material [put]
func (l LessonHandlerImpl) AddLessonMaterial(w http.ResponseWriter, r *http.Request) {
	payload := materials.AddLessonMaterialsRequestPayload{}

	// Validate payload
	if r.Header.Get("Content-Type") != "application/json" {
		payload := l.WrapperUtil.ErrorResponseWrap("this service only receive json input", nil)
		l.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	if err := l.HttpUtil.ParseJson(r, &payload); err != nil {
		payload := l.WrapperUtil.ErrorResponseWrap("invalid json input", err.Error())
		l.HttpUtil.WriteJson(w, http.StatusUnprocessableEntity, payload)
		return
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			payload := l.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
			l.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			return
		}

		errPayload := web.NewResponseErrorFromValidator(err.(validator.ValidationErrors))
		payload := l.WrapperUtil.ErrorResponseWrap(errPayload.Error(), errPayload)
		l.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	// Confirm Valid Website Token
	validateTokenHeader := r.Header.Get("Authorization")

	if validateTokenHeader == "" {
		payload := l.WrapperUtil.ErrorResponseWrap("token is required", nil)
		l.HttpUtil.WriteJson(w, http.StatusUnauthorized, payload)
		return
	}

	token := strings.Split(validateTokenHeader, " ")

	if len(token) != 2 {
		payload := l.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		l.HttpUtil.WriteJson(w, http.StatusUnauthorized, payload)
		return
	}

	if token[0] != "Bearer" {
		payload := l.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		l.HttpUtil.WriteJson(w, http.StatusUnauthorized, payload)
		return
	}

	payload.AddLessonMaterialsToken = token[1]
	err := l.LessonService.AddLessonMaterial(payload)

	if err != nil {
		if errData, ok := err.(web.ResponseError); ok {
			payload := l.WrapperUtil.ErrorResponseWrap(errData.Error(), errData)
			l.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			return
		}

		l.Logger.Error(
			fmt.Sprintf("[LESSON] some error happened when validating URL: %s", err.Error()),
		)
		payload := l.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		l.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	responsePayload := l.WrapperUtil.SuccessResponseWrap(nil)
	l.HttpUtil.WriteSuccessJson(w, responsePayload)
}
