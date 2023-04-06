package quiz

import "net/http"

type QuizHandler interface {
	GetAllQuizes(w http.ResponseWriter, r *http.Request)
	GetQuizDetail(w http.ResponseWriter, r *http.Request)
}
