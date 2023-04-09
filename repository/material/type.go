package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
)

type MaterialRepository interface {
	New(courseId string, creatorEmail string, name string) (uuid.UUID, error)
	Delete(id uuid.UUID) error
	Get(materialId uuid.UUID) (*material.Material, error)
	GetAll(courseId string) ([]material.Material, error)

	IsUserContributor(id uuid.UUID, email string) (bool, error)

	NewWithTransaction(tx transaction.Transaction, courseId string, creatorEmail string, name string) (uuid.UUID, error)
	DeleteWithTransaction(tx transaction.Transaction, id uuid.UUID) error
	GetAllWithTransaction(tx transaction.Transaction, courseId string) ([]material.Material, error)
}

type MaterialContentRepository interface {
	IsUserContributor(id uuid.UUID, email string) (bool, error)
	New(materialId uuid.UUID, materialType material.MaterialType, link string) (uuid.UUID, error)
	GetAll(materialId uuid.UUID) ([]material.Content, error)
	Delete(contentId uuid.UUID) error
	UpdateLink(contentId uuid.UUID, link string) error

	NewWithTransaction(tx transaction.Transaction, materialId uuid.UUID, materialType material.MaterialType, link string) (uuid.UUID, error)
	GetAllWithTransaction(tx transaction.Transaction, materialId uuid.UUID) ([]material.Content, error)
	DeleteWithTransaction(tx transaction.Transaction, contentId uuid.UUID) error
	UpdateLinkWithTransaction(tx transaction.Transaction, contentId uuid.UUID, link string) error
}
