package faculty

import "github.com/google/uuid"

// AddFaculty Request Payload
//	@Description	Information that should be available when you add a faculty
type AddFacultyRequestPayload struct {
	// Web Token that was appended to the link
	AddFacultyToken string `json:"faculty_token"`

	// Faculty Name
	Name string `json:"name" validate:"required"`

	// Faculty Name Abbreviation
	Abbreviation string `json:"abbreviation" validate:"required"`
}

// UpdateFaculty Request Payload
//	@Description	Information that should be available when you update a faculty
type UpdateFacultyRequestPayload struct {
	// Web Token that was appended to the link
	UpdateFacultyToken string

	// Faculty ID, Provided by Query
	ID uuid.UUID

	// Faculty Name
	Name string `json:"name" validate:"required"`

	// Faculty Name Abbreviation
	Abbreviation string `json:"abbreviation" validate:"required"`
}