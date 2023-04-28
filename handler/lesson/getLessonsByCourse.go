package lesson

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson"
)

// Index godoc
//
//	@Summary		Get lesson by course ID
//	@Description	Retrieve lesson data by course ID
//	@Tags			lesson
//	@Produce		json
//	@Param			id	path		string				true	"Course ID"
//	@Success		200	{object}	web.BaseResponse	"Success"
//	@Failure		400	{object}	web.BaseResponse	"Invalid ID provided in request path"
//	@Failure		401	{object}	web.BaseResponse	"Unauthorized"
//	@Failure		500	{object}	web.BaseResponse	"Internal Server Error"
//	@Router			/lesson/course/{id} [get]
func (l LessonHandlerImpl) GetLessonsByCourse(w http.ResponseWriter, r *http.Request) {
	payload := lesson.GetByStringRequestPayload{}

	payload.ID = chi.URLParam(r, "id")
	packet, err := l.LessonService.GetLessons(payload)

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