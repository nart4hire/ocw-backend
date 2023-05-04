package lesson

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/lesson"
)

type LessonRoutes struct {
	lesson.LessonHandler
}

func (l LessonRoutes) Register(r chi.Router) {
	r.Route("/lesson", func(r chi.Router) {
		r.Get("/{id}", l.LessonHandler.GetLesson)
		r.Get("/course/{id}", l.LessonHandler.GetLessonsByCourse)
		r.Get("/material/{id}", l.LessonHandler.GetLessonMaterial)
		r.Get("/material/lesson/{id}", l.LessonHandler.GetLessonMaterialsByLesson)

		r.Put("/", l.LessonHandler.AddLesson)
		r.Put("/material", l.LessonHandler.AddLessonMaterial)

		r.Post("/{id}", l.LessonHandler.UpdateLesson)
		r.Post("/material/{id}", l.LessonHandler.UpdateLessonMaterial)

		r.Delete("/{id}", l.LessonHandler.DeleteLesson)
		r.Delete("/material/{id}", l.LessonHandler.DeleteLessonMaterial)
	})
}
