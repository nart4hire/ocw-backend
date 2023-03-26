package course

import "net/http"

type CourseHandler interface {
	// Get
	GetCourses(w http.ResponseWriter, r *http.Request)
	GetCourse(w http.ResponseWriter, r *http.Request)
	GetMajors(w http.ResponseWriter, r *http.Request)
	GetMajor(w http.ResponseWriter, r *http.Request)
	GetCoursesByMajor(w http.ResponseWriter, r *http.Request)
	GetFaculty(w http.ResponseWriter, r *http.Request)
	GetFaculties(w http.ResponseWriter, r *http.Request)
	GetCoursesByFaculty(w http.ResponseWriter, r *http.Request)
	GetMajorsByFaculty(w http.ResponseWriter, r *http.Request)

	// Add (Put)
	AddCourse(w http.ResponseWriter, r *http.Request)
	AddMajor(w http.ResponseWriter, r *http.Request)
	AddFaculty(w http.ResponseWriter, r *http.Request)

	// Update
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	UpdateMajor(w http.ResponseWriter, r *http.Request)
	UpdateFaculty(w http.ResponseWriter, r *http.Request)

	// Delete
	DeleteCourse(w http.ResponseWriter, r *http.Request)
}
