package lesson

import (
	domLesson "gitlab.informatika.org/ocw/ocw-backend/model/domain/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson/materials"
)

type LessonService interface {
	GetLesson(payload lesson.GetByUUIDRequestPayload) (*domLesson.Lesson, error)
	GetLessons(payload lesson.GetByStringRequestPayload) ([]domLesson.Lesson, error)
	AddLesson(payload lesson.AddLessonRequestPayload) error
	UpdateLesson(payload lesson.UpdateLessonRequestPayload) error
	DeleteLesson(payload lesson.DeleteByUUIDRequestPayload) error
	GetLessonMaterial(payload materials.GetByUUIDRequestPayload) (*domLesson.LessonMaterials, error)
	GetLessonMaterials(payload materials.GetByUUIDRequestPayload) ([]domLesson.LessonMaterials, error)
	AddLessonMaterial(payload materials.AddLessonMaterialsRequestPayload) error
	UpdateLessonMaterial(payload materials.UpdateLessonMaterialsRequestPayload) error
	DeleteLessonMaterial(payload materials.DeleteByUUIDRequestPayload) error
}