package quiz

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
)

// Index godoc
//
//	@Tags					quiz
//	@Summary			Delete Quiz
//	@Description	Delete Quiz
//	@Produce			json
//	@Accept				json
//	@Param				id path string true "Quiz id" Format(uuid)
//	@Success			200	{object}	web.BaseResponse
//	@Router				/quiz/{id} [delete]
func (m QuizHandlerImpl) DeleteQuiz(w http.ResponseWriter, r *http.Request) {
	payload := quiz.DeleteRequestPayload{}
	quizId := chi.URLParam(r, "id")

	if quizId == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("quiz id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	id, err := uuid.Parse(quizId)

	if err != nil {
		// invalid uuid
		payload := m.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	// Confirm Valid Website Token
	validateTokenHeader := r.Header.Get("Authorization")

	if validateTokenHeader == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("token is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
		return
	}

	token := strings.Split(validateTokenHeader, " ")

	if len(token) != 2 {
		payload := m.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	if token[0] != "Bearer" {
		payload := m.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	payload.DeleteToken = token[1]
	payload.ID = id

	err = m.QuizService.DeleteQuiz(payload)

	if err != nil {
		respErr, ok := err.(web.ResponseError)
		if ok {
			payload := m.WrapperUtil.ErrorResponseWrap(respErr.Error(), respErr)

			m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		} else {
			payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
			m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		}
		return
	}

	responsePayload := m.WrapperUtil.SuccessResponseWrap(nil)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)

}