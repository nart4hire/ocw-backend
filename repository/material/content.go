package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gorm.io/gorm"
)

type MaterialContentRepositoryImpl struct {
	db *gorm.DB
}

func NewMaterialContent(
	db db.Database,
) *MaterialContentRepositoryImpl {
	return &MaterialContentRepositoryImpl{db.Connect()}
}

func (m MaterialContentRepositoryImpl) New(
	materialId uuid.UUID,
	materialType material.MaterialType,
	link string) (uuid.UUID, error) {
	contentData := material.Content{
		MaterialId: materialId,
		Type:       materialType,
		Link:       link,
	}

	if err := m.db.Create(&contentData).Error; err != nil {
		return uuid.Nil, err
	}

	return contentData.Id, nil
}

func (m MaterialContentRepositoryImpl) GetAll(materialId uuid.UUID) ([]material.Content, error) {
	result := []material.Content{}
	err := m.db.Where("material_id = ?", materialId).Find(&result).Error

	return result, err
}

func (m MaterialContentRepositoryImpl) Delete(contentId uuid.UUID) error {
	return m.db.Where("id = ?", contentId).Delete(&material.Content{}).Error
}

func (m MaterialContentRepositoryImpl) UpdateLink(contentId uuid.UUID, link string) error {
	return m.db.Where("id = ?", contentId).Update("link", link).Error
}
