package material

import (
	"github.com/google/uuid"
	materialDomain "gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/repository/material"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
)

type MaterialContentServiceImpl struct {
	transaction.TransactionBuilder
	material.MaterialContentRepository
}

func (m MaterialContentServiceImpl) AddContent(materialId uuid.UUID, user user.User, contents []materialDomain.Content) error {
	isSuccess := false
	tx := m.Build()

	tx.Begin()
	defer tx.Auto(&isSuccess)

	// TODO : Check user aman ga nambah konten

	for _, content := range contents {
		_, err := m.MaterialContentRepository.NewWithTransaction(tx, materialId, content.Type, content.Link)

		if err != nil {
			return err
		}
	}

	isSuccess = true
	return nil
}

func (m MaterialContentServiceImpl) DeleteContent(
	materialId uuid.UUID, user user.User, contentId uuid.UUID,
) error {
	// TODO: check user aman ga delete konten
	return m.MaterialContentRepository.Delete(contentId)
}

func (m MaterialContentServiceImpl) UpdateContentLink(materialId uuid.UUID, user user.User, contentId uuid.UUID, link string) error {
	// TODO: Check user aman ga update link
	return m.MaterialContentRepository.UpdateLink(contentId, link)
}
