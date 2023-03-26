package get

import "github.com/google/uuid"

// GetCourse Request Payload
// @Description Information that should be available when you get using course id (string)
type GetByStringRequestPayload struct {
	// Course ID, provided by query
	ID string
}

// GetCourse Request Payload
// @Description Information that should be available when you get using major/faculty id (string)
type GetByUUIDRequestPayload struct {
	// Major/Faculty ID, provided by query
	ID uuid.UUID
}