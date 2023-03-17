package delete

// DeleteCourse Request Payload
// @Description Information that should be available when you delete using course id (string)
type DeleteByStringRequestPayload struct {
	// Web Token that was appended to the link
	DeleteCourseToken string

	// Course ID, provided by query
	ID string
}
