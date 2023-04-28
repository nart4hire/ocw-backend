package lesson

import (
	"fmt"
	"net/http"
	
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson/materials"
)

// Index godoc
//
//	@Summary		Get lesson materials by lesson ID
//	@Description	Retrieve lesson materials data by lesson ID
//	@Tags			lesson
//	@Produce		json
//	@Param			id	path		string				true	"Lesson ID (UUID)"
//	@Success		200	{object}	web.BaseResponse	"Success"
//	@Failure		400	{object}	web.BaseResponse	"Invalid UUID provided in request path"
//	@Failure		401	{object}	web.BaseResponse	"Unauthorized"
//	@Failure		500	{object}	web.BaseResponse	"Internal Server Error"
//	@Router			/lesson/material/lesson/{id} [get]
func (l LessonHandlerImpl) GetLessonMaterialsByLesson(w http.ResponseWriter, r *http.Request) {
	payload := materials.GetByUUIDRequestPayload{}
	id, err := uuid.Parse(chi.URLParam(r, "id"))

	if err != nil {
		payload := l.WrapperUtil.ErrorResponseWrap("invalid id", nil)
		l.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}
	
	payload.ID = id
	packet, err := l.LessonService.GetLessonMaterials(payload)

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

	responsePayload := l.WrapperUtil.SuccessResponseWrap(packet)
	l.HttpUtil.WriteSuccessJson(w, responsePayload)
}