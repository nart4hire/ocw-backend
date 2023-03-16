package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type MaterialService interface {
	Create(courseId string, user user.User, contents []material.Content) (uuid.UUID, error)
	Delete(materialId uuid.UUID, user user.User) error
}

type MaterialContentService interface {
	AddContent(materialId uuid.UUID, user user.User, contents []material.Content) error
	DeleteContent(materialId uuid.UUID, user user.User, contentId uuid.UUID) error
	UpdateContentLink(materialId uuid.UUID, user user.User, contentId uuid.UUID, link string) error
}
