package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gorm.io/gorm"
)

type MaterialRepositoryImpl struct {
	db *gorm.DB
}

func NewMaterial(
	db db.Database,
) *MaterialRepositoryImpl {
	return &MaterialRepositoryImpl{db.Connect()}
}

func (m MaterialRepositoryImpl) New(courseId string, creatorEmail string) (uuid.UUID, error) {
	materialData := &material.Material{
		CourseId:     courseId,
		CreatorEmail: creatorEmail,
	}

	err := m.db.Create(materialData).Error

	if err != nil {
		return uuid.Nil, err
	}

	return materialData.Id, nil
}

func (m MaterialRepositoryImpl) Delete(id uuid.UUID) error {
	return m.db.Where("id = ?", id).Delete(&material.Material{}).Error
}

func (m MaterialRepositoryImpl) GetAll(courseId string) ([]material.Material, error) {
	result := []material.Material{}
	err := m.db.Joins("Contents").Where("CourseId = ?", courseId).Find(&result).Error

	return result, err
}
