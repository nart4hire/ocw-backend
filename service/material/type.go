package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
)

type MaterialService interface {
	Create(courseId string, contents []material.Content) (uuid.UUID, error)
	AddContent()
}
