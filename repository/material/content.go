package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gitlab.informatika.org/ocw/ocw-backend/repository/course"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
	"gorm.io/gorm"
)

type MaterialContentRepositoryImpl struct {
	builder transaction.TransactionBuilder
	course.CourseRepository
	MaterialRepository
	db *gorm.DB
}

func NewMaterialContent(
	builder transaction.TransactionBuilder,
	course course.CourseRepository,
	database db.Database,
	material MaterialRepository,
) *MaterialContentRepositoryImpl {
	return &MaterialContentRepositoryImpl{builder, course, material, database.Connect()}
}

func (m MaterialContentRepositoryImpl) IsUserContributor(id uuid.UUID, email string) (bool, error) {
	result := &material.Content{}
	if err := m.db.Where("id = ?", id).Find(result).Error; err != nil {
		return false, err
	}

	return m.MaterialRepository.IsUserContributor(result.MaterialID, email)
}

func (m MaterialContentRepositoryImpl) New(
	materialId uuid.UUID,
	materialType material.MaterialType,
	link string,
) (uuid.UUID, error) {
	return m.NewWithTransaction(
		m.builder.Build(),
		materialId,
		materialType,
		link,
	)
}

func (m MaterialContentRepositoryImpl) NewWithTransaction(
	tx transaction.Transaction,
	materialId uuid.UUID,
	materialType material.MaterialType,
	link string) (uuid.UUID, error) {
	contentData := material.Content{
		Id:         uuid.New(),
		MaterialID: materialId,
		Type:       materialType,
		Link:       link,
	}

	if err := tx.GetTransaction().Create(&contentData).Error; err != nil {
		return uuid.Nil, err
	}

	return contentData.Id, nil
}

func (m MaterialContentRepositoryImpl) GetAll(materialId uuid.UUID) ([]material.Content, error) {
	return m.GetAllWithTransaction(m.builder.Build(), materialId)
}
func (m MaterialContentRepositoryImpl) GetAllWithTransaction(tx transaction.Transaction, materialId uuid.UUID) ([]material.Content, error) {
	result := []material.Content{}
	err := tx.GetTransaction().Where("material_id = ?", materialId).Find(&result).Error

	return result, err
}

func (m MaterialContentRepositoryImpl) Delete(contentId uuid.UUID) error {
	return m.DeleteWithTransaction(m.builder.Build(), contentId)
}

func (m MaterialContentRepositoryImpl) DeleteWithTransaction(tx transaction.Transaction, contentId uuid.UUID) error {
	return tx.GetTransaction().Where("id = ?", contentId).Delete(&material.Content{}).Error
}

func (m MaterialContentRepositoryImpl) UpdateLink(contentId uuid.UUID, link string) error {
	return m.UpdateLinkWithTransaction(m.builder.Build(), contentId, link)
}

func (m MaterialContentRepositoryImpl) UpdateLinkWithTransaction(tx transaction.Transaction, contentId uuid.UUID, link string) error {
	return tx.GetTransaction().Where("id = ?", contentId).Update("link", link).Error
}
