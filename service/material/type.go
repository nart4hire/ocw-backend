package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
)

type MaterialService interface {
	Create(courseId string, email string, name string, week int) (uuid.UUID, error)
	Delete(materialId uuid.UUID, email string) error
	Get(courseId string) ([]material.Material, error)
	GetById(materialId uuid.UUID) (*material.Material, error)
}

type MaterialContentService interface {
	AddContent(materialId uuid.UUID, email string, content material.Content) (string, error)
	DeleteContent(materialId uuid.UUID, email string, contentId uuid.UUID) error
}
