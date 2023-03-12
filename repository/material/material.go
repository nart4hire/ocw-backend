package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
)

type MaterialRepositoryImpl struct {
	builder transaction.TransactionBuilder
}

func NewMaterial(
	builder transaction.TransactionBuilder,
) *MaterialRepositoryImpl {
	return &MaterialRepositoryImpl{builder}
}

func (m MaterialRepositoryImpl) New(courseId string, creatorEmail string) (uuid.UUID, error) {
	return m.NewWithTransaction(m.builder.Build(), courseId, creatorEmail)
}

func (m MaterialRepositoryImpl) NewWithTransaction(tx transaction.Transaction, courseId string, creatorEmail string) (uuid.UUID, error) {
	materialData := &material.Material{
		CourseId:     courseId,
		CreatorEmail: creatorEmail,
	}

	err := tx.GetTransaction().Create(materialData).Error

	if err != nil {
		return uuid.Nil, err
	}

	return materialData.Id, nil
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
	err := tx.GetTransaction().Joins("Contents").Where("CourseId = ?", courseId).Find(&result).Error

	return result, err
}
