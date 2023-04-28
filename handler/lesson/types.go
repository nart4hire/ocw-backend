package lesson

import "net/http"

type LessonHandler interface {
	// Get
	GetLesson(w http.ResponseWriter, r *http.Request)
	GetLessonsByCourse(w http.ResponseWriter, r *http.Request)
	GetLessonMaterial(w http.ResponseWriter, r *http.Request)
	GetLessonMaterialsByLesson(w http.ResponseWriter, r *http.Request)

	// Add (Put)
	AddLesson(w http.ResponseWriter, r *http.Request)
	AddLessonMaterial(w http.ResponseWriter, r *http.Request)

	// Update
	UpdateLesson(w http.ResponseWriter, r *http.Request)
	UpdateLessonMaterial(w http.ResponseWriter, r *http.Request)

	// Delete
	DeleteLesson(w http.ResponseWriter, r *http.Request)
	DeleteLessonMaterial(w http.ResponseWriter, r *http.Request)
}
