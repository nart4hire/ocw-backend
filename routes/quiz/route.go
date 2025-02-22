package quiz

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type QuizRoutes struct {
	quiz.QuizHandler
	*guard.GuardBuilder
}

func (q QuizRoutes) Register(r chi.Router) {
	r.Get("/course/{id}/quiz", q.QuizHandler.GetAllQuizes)

	guard := q.GuardBuilder.Build(
		user.Student,
		user.Contributor,
		user.Admin,
	)

	r.Route("/quiz/{id}/take", func(r chi.Router) {
		r.Use(guard)
		r.Post("/", q.QuizHandler.TakeQuiz)
	})

	r.Route("/quiz/{id}/finish", func(r chi.Router) {
		r.Use(guard)
		r.Post("/", q.QuizHandler.FinishQuiz)
	})

	r.Route("/quiz/{id}/solution", func(r chi.Router) {
		r.Use(guard)
		r.Get("/", q.QuizHandler.GetQuizSolution)
	})

	r.Route("/quiz/{id}", func(r chi.Router) {
		r.Get("/", q.QuizHandler.GetQuizDetail)
		r.Route("/", func(r chi.Router) {
			r.Use(guard)
			r.Patch("/", q.QuizHandler.UpdateQuiz)
			r.Delete("/", q.QuizHandler.DeleteQuiz)
		})
	})
	
	r.Route("/quiz", func(r chi.Router) {
		r.Use(guard)
		r.Put("/", q.QuizHandler.NewQuiz)
	})

	r.Route("/quiz/link/{id}", func(r chi.Router) {
		r.Use(guard)
		r.Get("/", q.QuizHandler.GetQuizLink)
	})
}
