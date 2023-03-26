package course

import (
	"github.com/go-chi/chi/v5"
	"gitlab.informatika.org/ocw/ocw-backend/handler/course"
)

type CourseRoutes struct {
	course.CourseHandler
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
	})
}
