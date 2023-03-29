package course

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/course"
	"gitlab.informatika.org/ocw/ocw-backend/handler/material"
	"gitlab.informatika.org/ocw/ocw-backend/middleware/guard"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type CourseRoutes struct {
	course.CourseHandler
	material.MaterialHandler
	*guard.GuardBuilder
}

func (c CourseRoutes) Register(r chi.Router) {
	r.Route("/course", func(r chi.Router) {
		// Get
		r.Get("/", c.CourseHandler.GetCourses)
		r.Get("/{id}", c.CourseHandler.GetCourse)
		r.Get("/faculty", c.CourseHandler.GetFaculties)
		r.Get("/faculty/{id}", c.CourseHandler.GetFaculty)
		r.Get("/faculty/courses/{id}", c.CourseHandler.GetCoursesByFaculty)
		r.Get("/faculty/majors/{id}", c.CourseHandler.GetMajorsByFaculty)
		r.Get("/major", c.CourseHandler.GetMajors)
		r.Get("/major/{id}", c.CourseHandler.GetMajor)
		r.Get("/major/courses/{id}", c.CourseHandler.GetCoursesByMajor)

		// Add
		r.Put("/", c.CourseHandler.AddCourse)
		r.Put("/faculty", c.CourseHandler.AddFaculty)
		r.Put("/major", c.CourseHandler.AddMajor)

		// Update
		r.Patch("/{id}", c.CourseHandler.UpdateCourse)
		r.Patch("/faculty/{id}", c.CourseHandler.UpdateFaculty)
		r.Patch("/major/{id}", c.CourseHandler.UpdateMajor)

		// Delete
		r.Delete("/{id}", c.CourseHandler.DeleteCourse)
		r.Get("/{id}/materials", c.MaterialHandler.GetMaterial)
	})

	r.Route("/course/{id}/material", func(r chi.Router) {
		r.Use(c.BuildSimple(user.Contributor))
		r.Post("/", c.MaterialHandler.CreateMaterial)
	})
}
