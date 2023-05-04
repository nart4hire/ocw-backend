package lesson

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/lesson"
	"github.com/google/uuid"
)

type LessonRepository interface {
	GetLesson(id uuid.UUID) (*lesson.Lesson, error)
	GetLessons(courseId string) ([]lesson.Lesson, error)
	AddLesson(lesson lesson.Lesson) error
	UpdateLesson(lesson lesson.Lesson) error
	DeleteLesson(id uuid.UUID) error
}

type LessonMaterialsRepository interface {
	GetLessonMaterial(id uuid.UUID) (*lesson.LessonMaterials, error)
	GetLessonMaterials(lessonId uuid.UUID) ([]lesson.LessonMaterials, error)
	AddLessonMaterial(lessonMaterial lesson.LessonMaterials) error
	UpdateLessonMaterial(lessonMaterial lesson.LessonMaterials) error
	DeleteLessonMaterial(id uuid.UUID) error
}