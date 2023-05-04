package quiz

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/quiz"
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

			m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		} else {
			payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
			m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		}
		return
	}

	responsePayload := m.WrapperUtil.SuccessResponseWrap(result)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)

}

// Index godoc
//
//	@Tags					quiz
//	@Summary			Get Quiz Link
//	@Description	Get Quiz Link
//	@Produce			json
//	@Accept				json
//	@Param				id path string true "Quiz id" Format(uuid)
//	@Success			200	{object}	web.BaseResponse
//	@Router				/quiz/link/{id} [get]

func (m QuizHandlerImpl) GetQuizLink(w http.ResponseWriter, r *http.Request) {
	payload := quiz.GetRequestPayload{}
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

	payload.GetToken = token[1]
	payload.ID = id
	response, err := m.QuizService.GetQuizLink(payload)

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

	responsePayload := m.WrapperUtil.SuccessResponseWrap(response)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)
}
