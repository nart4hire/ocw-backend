package update

import "github.com/google/uuid"

// UpdateCourse Request Payload
// @Description Information that should be available when you add a course
type UpdateCourseRequestPayload struct {
	// Web Token that was appended to the link
	UpdateCourseToken string

	// Course ID, Provided by query
	ID string `json:"id"`

	// Course Name
	Name string `json:"name" validate:"required"`

	// Course Major Abbreviation
	MajAbbr string `json:"majabbr" validate:"required_without=MajorID"`

	// Major Id, will be set by the server
	MajorID uuid.UUID `json:"major_id"`

	// Course Description (Can be left empty)
	Description string `json:"description"`

	// Contributor Email
	Email string `json:"email" validate:"required,email" example:"someone@example.com"`

	// Course Name Abbreviation
	Abbreviation string `json:"abbreviation" validate:"required"`

	// Course Lecturer
	Lecturer string `json:"lecturer" validate:"required"`
}
