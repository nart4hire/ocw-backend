package quiz

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/quiz"
)

type QuizRoutes struct {
	quiz.QuizHandler
}

func (q QuizRoutes) Register(r chi.Router) {
	r.Get("/course/{id}/quiz", q.QuizHandler.GetAllQuizes)
	r.Get("/quiz/{id}", q.QuizHandler.GetQuizDetail)
}
