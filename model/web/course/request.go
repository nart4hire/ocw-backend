package course

import "github.com/google/uuid"

// AddCourse Request Payload
//	@Description	Information that should be available when you add a course
type AddCourseRequestPayload struct {
	// Web Token that was appended to the link
	AddCourseToken string

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
	Abbreviation string `json:"abbreviation"`
}

// DeleteCourse Request Payload
//	@Description	Information that should be available when you delete using course id (string)
type DeleteByStringRequestPayload struct {
	// Web Token that was appended to the link
	DeleteCourseToken string

	// Course ID, provided by query
	ID string `json:"-" validate:"required"`
}


// GetID Request Payload
//	@Description	Information that should be available when you get using course id (string)
type GetByStringRequestPayload struct {
	// Course ID, provided by query
	ID string `json:"-" validate:"required"`
}

// GetUUID Request Payload
//	@Description	Information that should be available when you get using major/faculty id (string)
type GetByUUIDRequestPayload struct {
	// Major/Faculty ID, provided by query
	ID uuid.UUID `json:"-" validate:"required"`
}

// UpdateCourse Request Payload
//	@Description	Information that should be available when you add a course
type UpdateCourseRequestPayload struct {
	// Web Token that was appended to the link
	UpdateCourseToken string

	// Course ID, Provided by query
	ID string `json:"-"`

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


