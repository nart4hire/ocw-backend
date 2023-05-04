package lesson

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gorm.io/gorm"
)

type LessonRepositoryImpl struct {
	db *gorm.DB
}

func NewLesson(
	db db.Database,
) *LessonRepositoryImpl {
	return &LessonRepositoryImpl{db.Connect()}
}

func (repo LessonRepositoryImpl) GetLesson(id uuid.UUID) (*lesson.Lesson, error) {
	result := &lesson.Lesson{}
	err := repo.db.First(result, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo LessonRepositoryImpl) GetLessons(courseId string) ([]lesson.Lesson, error) {
	var result []lesson.Lesson
	err := repo.db.Where("course_id = ?", courseId).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo LessonRepositoryImpl) AddLesson(lesson lesson.Lesson) error {
	return repo.db.Create(lesson).Error
}

func (repo LessonRepositoryImpl) UpdateLesson(lesson lesson.Lesson) error {
	return repo.db.Save(lesson).Error
}

func (repo LessonRepositoryImpl) DeleteLesson(id uuid.UUID) error {
	return repo.db.Delete(&lesson.Lesson{}, id).Error
}
