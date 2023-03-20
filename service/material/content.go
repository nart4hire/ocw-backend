package material

import (
	"errors"

	"github.com/google/uuid"
	materialDomain "gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/repository/material"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
	"gorm.io/gorm"
)

type MaterialContentServiceImpl struct {
	transaction.TransactionBuilder
	material.MaterialContentRepository
}

func (m MaterialContentServiceImpl) isMaterialContributor(materialId uuid.UUID, user user.User) error {
	_, err := m.MaterialContentRepository.IsUserContributor(materialId, user.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.NewResponseError("materials and user combination not found", "ERR_MATERIAL_USER_NOT_FOUND")
		}

		return err
	}

	return nil
}

func (m MaterialContentServiceImpl) AddContent(materialId uuid.UUID, user user.User, contents []materialDomain.Content) error {
	// TODO : Check user aman ga nambah konten
	if err := m.isMaterialContributor(materialId, user); err != nil {
		return err
	}

	isSuccess := false
	tx := m.Build()

	tx.Begin()
	defer tx.Auto(&isSuccess)

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
	if err := m.isMaterialContributor(materialId, user); err != nil {
		return err
	}

	return m.MaterialContentRepository.Delete(contentId)
}

func (m MaterialContentServiceImpl) UpdateContentLink(materialId uuid.UUID, user user.User, contentId uuid.UUID, link string) error {
	// TODO: Check user aman ga update link
	if err := m.isMaterialContributor(materialId, user); err != nil {
		return err
	}

	return m.MaterialContentRepository.UpdateLink(contentId, link)
}
