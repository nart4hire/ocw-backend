package quiz

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/quiz"
)

// Index godoc
//
//	@Tags					quiz
//	@Summary			New Quiz
//	@Description	New Quiz
//	@Produce			json
//	@Accept				json
//	@Param				id path string true "Quiz id" Format(uuid)
//	@Success			200	{object}	web.BaseResponse
//	@Router				/quiz [put]
func (m QuizHandlerImpl) NewQuiz(w http.ResponseWriter, r *http.Request) {
	payload := quiz.AddQuizRequestPayload{}

	// Validate payload
	if r.Header.Get("Content-Type") != "application/json" {
		payload := m.WrapperUtil.ErrorResponseWrap("this service only receive json input", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	if err := m.HttpUtil.ParseJson(r, &payload); err != nil {
		payload := m.WrapperUtil.ErrorResponseWrap("invalid json input", err.Error())
		m.HttpUtil.WriteJson(w, http.StatusUnprocessableEntity, payload)
		return
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			payload := m.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
			m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			return
		}

		errPayload := web.NewResponseErrorFromValidator(err.(validator.ValidationErrors))
		payload := m.WrapperUtil.ErrorResponseWrap(errPayload.Error(), errPayload)
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

	payload.AddQuizToken = token[1]

	response, err := m.QuizService.NewQuiz(payload)

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