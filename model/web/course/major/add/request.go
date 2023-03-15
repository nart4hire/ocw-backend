package add

import "github.com/google/uuid"

// AddMajor Request Payload
// @Description Information that should be available when you add a major
type AddMajorRequestPayload struct {
	// Major Name
	Name string `json:"name" validate:"required"`

	// Major Faculty Abbreviation
	FacAbbr string `json:"facabbr" validate:"required_without=FacultyID"`

	// Faculty Id, will be set by the server
	FacultyID uuid.UUID `json:"faculty_id"`

	// Major Name Abbreviation
	Abbreviation string `json:"abbreviation" validate:"required"`
}
