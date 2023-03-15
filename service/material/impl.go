package material

import (
	"github.com/google/uuid"
	materialRepo "gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/repository/material"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
)

type MaterialServiceImpl struct {
	transaction.TransactionBuilder
	material.MaterialContentRepository
	material.MaterialRepository
}

func (m MaterialServiceImpl) Create(courseId string, user user.User, contents []materialRepo.Content) (uuid.UUID, error) {
	isSuccess := false
	tx := m.TransactionBuilder.Build()

	tx.Begin()
	defer tx.Auto(&isSuccess)

	id, err := m.MaterialRepository.NewWithTransaction(tx, courseId, user.Email)

	if err != nil {
		return uuid.Nil, err
	}

	for _, content := range contents {
		_, err = m.MaterialContentRepository.NewWithTransaction(tx, id, content.Type, content.Link)

		if err != nil {
			return uuid.Nil, err
		}
	}

	isSuccess = true
	return id, err
}

func (m MaterialServiceImpl) Delete(materialId uuid.UUID, user user.User) error {
	// TODO: Pengecekan user apakah kontributor course bukan
	return m.MaterialRepository.Delete(materialId)
}
