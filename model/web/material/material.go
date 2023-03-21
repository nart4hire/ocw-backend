package material

import "github.com/google/uuid"

type CreateMaterialRequest struct {
	CourseId string `json:"course_id"`
}

type CreateMaterialResponse struct {
	MaterialId uuid.UUID `json:"material_id"`
}

type DeleteMaterialRequest struct {
	MaterialId uuid.UUID `json:"material_id"`
}
