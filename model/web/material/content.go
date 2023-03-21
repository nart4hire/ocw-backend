package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
)

type NewContentRequest struct {
	Type       material.MaterialType `json:"type"`
	Link       string                `json:"link"`
	MaterialId uuid.UUID             `json:"-"`
}

type NewContentResponse struct {
	UploadLink string `json:"upload_link"`
}

type DeleteContentRequest struct {
	ContentId  uuid.UUID `json:"content_id"`
	MaterialId uuid.UUID `json:"-"`
}
