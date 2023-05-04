package lesson

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gorm.io/gorm"
)

type LessonMaterialsRepositoryImpl struct {
	db *gorm.DB
}

func NewLessonMaterials(
	db db.Database,
) *LessonMaterialsRepositoryImpl {
	return &LessonMaterialsRepositoryImpl{db.Connect()}
}

func (repo LessonMaterialsRepositoryImpl) GetLessonMaterial(id uuid.UUID) (*lesson.LessonMaterials, error) {
	result := &lesson.LessonMaterials{}
	err := repo.db.First(result, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo LessonMaterialsRepositoryImpl) GetLessonMaterials(lessonId uuid.UUID) ([]lesson.LessonMaterials, error) {
	var result []lesson.LessonMaterials
	err := repo.db.Where("lesson_id = ?", lessonId).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo LessonMaterialsRepositoryImpl) AddLessonMaterial(lessonMaterial lesson.LessonMaterials) error {
	return repo.db.Create(lessonMaterial).Error
}

func (repo LessonMaterialsRepositoryImpl) UpdateLessonMaterial(lessonMaterial lesson.LessonMaterials) error {
	return repo.db.Save(lessonMaterial).Error
}

func (repo LessonMaterialsRepositoryImpl) DeleteLessonMaterial(id uuid.UUID) error {
	return repo.db.Delete(&lesson.LessonMaterials{}, id).Error
}
