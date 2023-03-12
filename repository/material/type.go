package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
)

type MaterialRepository interface {
	New(courseId string, creatorEmail string) (uuid.UUID, error)
	Delete(id uuid.UUID) error
	GetAll(courseId string) ([]material.Material, error)
}

type MaterialContentRepository interface {
	New(materialId uuid.UUID, materialType material.MaterialType, link string) (uuid.UUID, error)
	GetAll(materialId uuid.UUID) ([]material.Content, error)
	Delete(contentId uuid.UUID) error
	UpdateLink(contentId uuid.UUID, link string) error
}
