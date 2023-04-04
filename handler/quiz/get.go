package quiz

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
)

// Index godoc
//
//	@Tags					quiz
//	@Summary			Get Quiz Detail
//	@Description	Get Quiz Detail
//	@Produce			json
//	@Accept				json
//	@Param				id path string true "Quiz id" Format(uuid)
//	@Success			200	{object}	web.BaseResponse{data=quiz.Quiz}
//	@Router				/quiz/{id} [get]
func (m QuizHandlerImpl) GetQuizDetail(w http.ResponseWriter, r *http.Request) {
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

	result, err := m.QuizService.GetQuizDetail(id)

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

	responsePayload := m.WrapperUtil.SuccessResponseWrap(result)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)

}
