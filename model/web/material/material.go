package material

import "github.com/google/uuid"

type CreateMaterialRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateMaterialResponse struct {
	MaterialId uuid.UUID `json:"material_id"`
}

type DeleteMaterialRequest struct {
	MaterialId uuid.UUID `json:"material_id"`
}
