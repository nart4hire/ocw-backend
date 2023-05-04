package quiz

import (
	"github.com/google/uuid"
)

// AddQuiz Request Payload
//
//	@Description	Information that should be available when you add a quiz
type AddQuizRequestPayload struct {
	// Web Token that was appended to the link
	AddQuizToken string

	// Quiz Name
	Name string `json:"name" validate:"required"`

	// Course ID
	CourseID string `json:"course_id" validate:"required"`
}

// UpdateQuiz Request Payload
//
//	@Description	Information that should be available when you update a quiz
type UpdateQuizRequestPayload struct {
	// Web Token that was appended to the link
	UpdateQuizToken string

	// Quiz ID, Set by param
	ID uuid.UUID `json:"id"`
}

// DeleteQuiz Request Payload
//
//	@Description	Information that should be available when you delete using uuid
type DeleteRequestPayload struct {
	// Web Token that was appended to the link
	DeleteToken string

	// Quiz ID, Set by param
	ID uuid.UUID
}

// GetUUID Request Payload
//
//	@Description	Information that should be available when you get using uuid
type GetRequestPayload struct {
	// Web Token that was appended to the link
	GetToken string

	// Quiz/Problem/Answer ID, provided by query
	ID uuid.UUID
}

// Link Response Payload
//
//	@Description	Information that you will get upon successful add/update request
type LinkResponse struct {
	UploadLink string `json:"upload_link"`
}

// Path Response Payload
//
//	@Description	Information that you will get upon successful get request
type PathResponse struct {
	Path string `json:"path"`
}