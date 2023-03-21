package material

import (
	"errors"

	"github.com/google/uuid"
	materialDomain "gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/repository/material"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
	"gorm.io/gorm"
)

type MaterialServiceImpl struct {
	transaction.TransactionBuilder
	material.MaterialContentRepository
	material.MaterialRepository
}

func (m MaterialServiceImpl) Get(courseId string) ([]materialDomain.Material, error) {
	materials, err := m.MaterialRepository.GetAll(courseId)
	return materials, err
}

func (m MaterialServiceImpl) Create(courseId string, email string) (uuid.UUID, error) {
	isSuccess := false
	tx := m.TransactionBuilder.Build()

	tx.Begin()
	defer tx.Auto(&isSuccess)

	id, err := m.MaterialRepository.NewWithTransaction(tx, courseId, email)

	if err != nil {
		return uuid.Nil, err
	}

	isSuccess = true
	return id, err
}

func (m MaterialServiceImpl) Delete(materialId uuid.UUID, email string) error {
	// TODO: Pengecekan user apakah kontributor course bukan
	_, err := m.MaterialRepository.IsUserContributor(materialId, email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.NewResponseError("User is not the owner of material", "NOT_OWNER")
		}

		return err
	}

	return m.MaterialRepository.Delete(materialId)
}
