package material

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	materialDomain "gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/provider/storage"
	"gitlab.informatika.org/ocw/ocw-backend/repository/material"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gorm.io/gorm"
)

type MaterialContentServiceImpl struct {
	transaction.TransactionBuilder
	material.MaterialContentRepository
	material.MaterialRepository
	storage.Storage
	logger.Logger
	*env.Environment
}

func (m MaterialContentServiceImpl) isMaterialContributor(materialId uuid.UUID, email string) error {
	_, err := m.MaterialRepository.IsUserContributor(materialId, email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.NewResponseError("materials and user combination not found", "NOT_OWNER")
		}

		return err
	}

	return nil
}

func (m MaterialContentServiceImpl) AddContent(materialId uuid.UUID, user string, content materialDomain.Content) (string, error) {
	// TODO : Check user aman ga nambah konten
	if err := m.isMaterialContributor(materialId, user); err != nil {
		return "", err
	}

	isSuccess := false
	tx := m.Build()

	tx.Begin()
	defer tx.Auto(&isSuccess)

	if content.Type == materialDomain.Handout {
		path := fmt.Sprintf("%s/%s.pdf", m.BucketMaterialBasePath, strings.ReplaceAll(uuid.New().String(), "-", ""))
		uploadLink, err := m.Storage.CreatePutSignedLink(context.Background(), path)

		if err != nil {
			m.Logger.Error("Some error happened when generate link")
			m.Logger.Error(err.Error())
			return "", err
		}

		_, err = m.MaterialContentRepository.NewWithTransaction(tx, materialId, content.Type, path)

		if err != nil {
			return "", err
		}

		isSuccess = true

		return uploadLink, nil
	} else {
		if content.Link == "" {
			return "", web.NewResponseError("content is empty", "ERR_CONTENT_LINK_EMPTY")
		}

		_, err := m.MaterialContentRepository.NewWithTransaction(tx, materialId, content.Type, content.Link)

		if err == nil {
			isSuccess = true
		}

		return "", err
	}
}

func (m MaterialContentServiceImpl) DeleteContent(
	materialId uuid.UUID, user string, contentId uuid.UUID,
) error {
	// TODO: check user aman ga delete konten
	if err := m.isMaterialContributor(materialId, user); err != nil {
		return err
	}

	return m.MaterialContentRepository.Delete(contentId)
}
