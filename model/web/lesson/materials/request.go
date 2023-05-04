package materials

import (
	"github.com/google/uuid"
)

// AddLessonMaterials Request Payload
//	@Description	Information that should be available when you add a lesson material
type AddLessonMaterialsRequestPayload struct {
	// Web Token that was appended to the link
	AddLessonMaterialsToken string

	// Lesson ID
	LessonID uuid.UUID `json:"lesson_id" validate:"required"`

	// Lesson Material Order
	Order int `json:"order" validate:"required"`

	// Lesson Material ID, optional
	MaterialID uuid.UUID `json:"material_id"`

	// Lesson Contents
	Contents string `json:"contents" validate:"required"`
}

// DeleteLessonMaterials Request Payload
//	@Description	Information that should be available when you delete using lesson material id (uuid)
type DeleteByUUIDRequestPayload struct {
	// Web Token that was appended to the link
	DeleteLessonMaterialsToken string

	// Lesson Material ID, provided by query
	ID uuid.UUID `json:"-" validate:"required"`
}

// UpdateLessonMaterials Request Payload
//	@Description	Information that should be available when you update a lesson material
type UpdateLessonMaterialsRequestPayload struct {
	// Web Token that was appended to the link
	UpdateLessonMaterialsToken string

	// Lesson Material ID, provided by query
	ID uuid.UUID `json:"-" validate:"required"`

	// Lesson ID
	LessonID uuid.UUID `json:"lesson_id" validate:"required"`

	// Lesson Material Order
	Order int `json:"order" validate:"required"`

	// Lesson Material ID, optional
	MaterialID uuid.UUID `json:"material_id"`

	// Lesson Contents
	Contents string `json:"contents" validate:"required"`
}

// GetUUID Request Payload
//	@Description	Information that should be available when you get using id or lesson id (uuid)
type GetByUUIDRequestPayload struct {
	// Lesson ID, provided by query
	ID uuid.UUID `json:"-" validate:"required"`
}