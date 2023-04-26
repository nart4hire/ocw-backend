package quiz

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	authToken "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/quiz"
)

// Index godoc
//
//	@Tags					quiz
//	@Summary			Take Quiz
//	@Description	Take a quiz
//	@Produce			json
//	@Accept				json
//	@Param				Authorization	header		string							true "Authenticate User (any role)"
//	@Param				id path string true "Quiz id" Format(uuid)
//	@Success			200	{object}	web.BaseResponse{data=quiz.QuizDetail}
//	@Router				/quiz/{id}/take [post]
func (m QuizHandlerImpl) TakeQuiz(w http.ResponseWriter, r *http.Request) {
	rawQuizId := chi.URLParam(r, "id")

	if rawQuizId == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("quiz id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	quizId, err := uuid.Parse(rawQuizId)

	if err != nil {
		payload := m.WrapperUtil.ErrorResponseWrap("quiz id is not valid", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	user, ok := r.Context().Value(guard.UserContext).(authToken.UserClaim)

	if !ok {
		m.Logger.Error("Context is not found")
		payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	detail, err := m.DoTakeQuiz(r.Context(), quizId, user.Email)

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

	responsePayload := m.WrapperUtil.SuccessResponseWrap(detail)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)
}

// Index godoc
//
//	@Tags					quiz
//	@Summary			Get Quiz Solution
//	@Description	Take a quiz
//	@Produce			json
//	@Accept				json
//	@Param				Authorization	header		string							true "Authenticate User (any role)"
//	@Param				id path string true "Quiz id" Format(uuid)
//	@Success			200	{object}	web.BaseResponse{data=quiz.QuizDetail}
//	@Router				/quiz/{id}/solution [get]
func (m QuizHandlerImpl) GetQuizSolution(w http.ResponseWriter, r *http.Request) {
	rawQuizId := chi.URLParam(r, "id")

	if rawQuizId == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("quiz id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	quizId, err := uuid.Parse(rawQuizId)

	if err != nil {
		payload := m.WrapperUtil.ErrorResponseWrap("quiz id is not valid", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	user, ok := r.Context().Value(guard.UserContext).(authToken.UserClaim)

	if !ok {
		m.Logger.Error("Context is not found")
		payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	detail, err := m.GetSolutionQuiz(r.Context(), quizId, user.Email)

	if err != nil {
		respErr, ok := err.(web.ResponseError)
		if ok {
			if respErr.Code == "ERR_NOT_ALLOWED" {
				payload := m.WrapperUtil.ErrorResponseWrap(respErr.Error(), respErr)
				m.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
			} else {
				payload := m.WrapperUtil.ErrorResponseWrap(respErr.Error(), respErr)
				m.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			}
		} else {
			payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
			m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		}
		return
	}

	responsePayload := m.WrapperUtil.SuccessResponseWrap(detail)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)
}

// Index godoc
//
//	@Tags					quiz
//	@Summary			Finish Quiz
//	@Description	Finish quiz session and get the score
//	@Produce			json
//	@Accept				json
//	@Param				Authorization	header		string										true "Authenticate User (any role)"
//	@Param			  data	body		quiz.FinishQuizPayload							true	"Quiz Finish payload"
//	@Param				id path string true "Quiz id" Format(uuid)
//	@Success			200	{object}	web.BaseResponse{data=quiz.QuizDetail}
//	@Router				/quiz/{id}/take [post]
func (m QuizHandlerImpl) FinishQuiz(w http.ResponseWriter, r *http.Request) {
	payload := quiz.FinishQuizPayload{}

	/* Get user */
	user, ok := r.Context().Value(guard.UserContext).(authToken.UserClaim)

	if !ok {
		m.Logger.Error("Context is not found")
		payload := m.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		m.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}
	/* Get user */

	/* Validate input */
	validate := validator.New()

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
	/* End of validate */

	/* Get quiz id */
	rawQuizId := chi.URLParam(r, "id")

	if rawQuizId == "" {
		payload := m.WrapperUtil.ErrorResponseWrap("quiz id is required", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	quizId, err := uuid.Parse(rawQuizId)

	if err != nil {
		payload := m.WrapperUtil.ErrorResponseWrap("quiz id is not valid", nil)
		m.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}
	/* end of get quiz id */

	res, err := m.DoFinishQuiz(r.Context(), quizId, user.Email, payload.Data)

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

	responsePayload := m.WrapperUtil.SuccessResponseWrap(res)
	m.HttpUtil.WriteSuccessJson(w, responsePayload)
}
