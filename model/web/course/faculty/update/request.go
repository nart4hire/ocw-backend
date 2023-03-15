package update

import "github.com/google/uuid"

// UpdateFaculty Request Payload
// @Description Information that should be available when you update a faculty
type UpdateFacultyRequestPayload struct {
	// Faculty ID, Provided by Query
	ID uuid.UUID

	// Faculty Name
	Name string `json:"name" validate:"required"`

	// Faculty Name Abbreviation
	Abbreviation string `json:"abbreviation" validate:"required"`
}
