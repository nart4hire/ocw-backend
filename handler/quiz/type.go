package quiz

import "net/http"

type QuizHandler interface {
	GetAllQuizes(w http.ResponseWriter, r *http.Request)
	GetQuizDetail(w http.ResponseWriter, r *http.Request)

	TakeQuiz(w http.ResponseWriter, r *http.Request)
	GetQuizSolution(w http.ResponseWriter, r *http.Request)
	FinishQuiz(w http.ResponseWriter, r *http.Request)
	NewQuiz(w http.ResponseWriter, r *http.Request)
	GetQuizLink(w http.ResponseWriter, r *http.Request)
	UpdateQuiz(w http.ResponseWriter, r *http.Request)
	DeleteQuiz(w http.ResponseWriter, r *http.Request)
}
