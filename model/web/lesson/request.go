package lesson

import "github.com/google/uuid"

// AddLesson Request Payload
//	@Description	Information that should be available when you add a lesson
type AddLessonRequestPayload struct {
	// Web Token that was appended to the link
	AddLessonToken string

	// Lesson Name
	Name string `json:"name" validate:"required"`

	// Course ID
	CourseID string `json:"course_id" validate:"required"`
	
	// Lesson Order
	Order int `json:"order" validate:"required"`

	// Lesson Description (Can be left empty)
	Description string `json:"description"`
}

// DeleteLesson Request Payload
//	@Description	Information that should be available when you delete using lesson id (uuid)
type DeleteByUUIDRequestPayload struct {
	// Web Token that was appended to the link
	DeleteLessonToken string

	// Lesson ID, provided by query
	ID uuid.UUID `json:"-" validate:"required"`
}

// UpdateLesson Request Payload
//	@Description	Information that should be available when you update a lesson
type UpdateLessonRequestPayload struct {
	// Web Token that was appended to the link
	UpdateLessonToken string

	// Lesson ID, provided by query
	ID uuid.UUID `json:"-" validate:"required"`

	// Lesson Name
	Name string `json:"name" validate:"required"`

	// Course ID
	CourseID string `json:"course_id" validate:"required"`

	// Lesson Order
	Order int `json:"order" validate:"required"`

	// Lesson Description (Can be left empty)
	Description string `json:"description"`
}

// GetUUID Request Payload
//	@Description	Information that should be available when you get using lesson id (uuid)
type GetByUUIDRequestPayload struct {
	// Lesson ID, provided by query
	ID uuid.UUID `json:"-" validate:"required"`
}

// GetID Request Payload
//	@Description	Information that should be available when you get using course id (string)
type GetByStringRequestPayload struct {
	// Course ID, provided by query
	ID string `json:"-" validate:"required"`
}
