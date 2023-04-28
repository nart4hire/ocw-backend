package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
	"gorm.io/gorm"
)

type MaterialRepositoryImpl struct {
	builder transaction.TransactionBuilder
	db      *gorm.DB
}

func NewMaterial(
	builder transaction.TransactionBuilder,
	db db.Database,
) *MaterialRepositoryImpl {
	return &MaterialRepositoryImpl{builder, db.Connect()}
}

func (m MaterialRepositoryImpl) Get(materialId uuid.UUID) (*material.Material, error) {
	res := &material.Material{}
	err := m.db.Preload("Contents").Where("id = ?", materialId).Find(res).Error
	return res, err
}

func (m MaterialRepositoryImpl) IsUserContributor(id uuid.UUID, email string) (bool, error) {
	err := m.db.Where("creator_email = ? AND id = ?", email, id).Find(&material.Material{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m MaterialRepositoryImpl) New(courseId string, creatorEmail string, name string, week int) (uuid.UUID, error) {
	return m.NewWithTransaction(m.builder.Build(), courseId, creatorEmail, name, week)
}

func (m MaterialRepositoryImpl) NewWithTransaction(tx transaction.Transaction, courseId string, creatorEmail string, name string, week int) (uuid.UUID, error) {
	materialData := &material.Material{
		ID:           uuid.New(),
		CourseId:     courseId,
		CreatorEmail: creatorEmail,
		Name:         name,
		Week:         week,
	}

	err := tx.GetTransaction().Create(materialData).Error

	if err != nil {
		return uuid.Nil, err
	}

	return materialData.ID, nil
}

func (m MaterialRepositoryImpl) Delete(id uuid.UUID) error {
	return m.DeleteWithTransaction(m.builder.Build(), id)
}

func (m MaterialRepositoryImpl) DeleteWithTransaction(tx transaction.Transaction, id uuid.UUID) error {
	return tx.GetTransaction().Where("id = ?", id).Delete(&material.Material{}).Error
}

func (m MaterialRepositoryImpl) GetAll(courseId string) ([]material.Material, error) {
	return m.GetAllWithTransaction(m.builder.Build(), courseId)
}

func (m MaterialRepositoryImpl) GetAllWithTransaction(tx transaction.Transaction, courseId string) ([]material.Material, error) {
	result := []material.Material{}
	trx := tx.GetTransaction()
	err := trx.
		Model(&material.Material{}).
		Preload("Contents").
		Where("course_id = ?", courseId).
		Find(&result).Error

	return result, err
}
