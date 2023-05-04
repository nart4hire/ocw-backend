package lesson

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson/materials"
)

// Index godoc
//
//	@Summary		Delete lesson material by id
//	@Description	Delete a lesson material with the specified ID
//  @Tags			lesson
//	@Accept			json
//	@Produce		json
//  @Param          id			  path        string                          true        "Lesson Material ID"
//	@Param			Authorization	header		string							true	"DeleteLessonMaterialToken"
//	@Success		200				{object}	web.BaseResponse				"Success"
//	@Failure		400				{object}	web.BaseResponse				"Bad Request"
//	@Failure		401				{object}	web.BaseResponse				"Unauthorized"
//	@Failure		403				{object}	web.BaseResponse				"Forbidden"
//	@Failure		422				{object}	web.BaseResponse				"Unprocessable Entity"
//	@Failure		500				{object}	web.BaseResponse				"Internal Server Error"
//	@Router			/lesson/material/{id} [delete]
func (l LessonHandlerImpl) DeleteLessonMaterial(w http.ResponseWriter, r *http.Request) {
	payload := materials.DeleteByUUIDRequestPayload{}
	validateTokenHeader := r.Header.Get("Authorization")

	if validateTokenHeader == "" {
		payload := l.WrapperUtil.ErrorResponseWrap("token is required", nil)
		l.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
		return
	}

	token := strings.Split(validateTokenHeader, " ")

	if len(token) != 2 {
		payload := l.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		l.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	if token[0] != "Bearer" {
		payload := l.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		l.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	id, err := uuid.Parse(chi.URLParam(r, "id"))

	if err != nil {
		payload := l.WrapperUtil.ErrorResponseWrap("invalid id", nil)
		l.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}


	payload.DeleteLessonMaterialsToken = token[1]
	payload.ID = id
	err = l.LessonService.DeleteLessonMaterial(payload)

	if err != nil {
		if errData, ok := err.(web.ResponseError); ok {
			payload := l.WrapperUtil.ErrorResponseWrap(errData.Error(), errData)
			l.HttpUtil.WriteJson(w, http.StatusUnauthorized, payload)
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
