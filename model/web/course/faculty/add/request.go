package add

// AddFaculty Request Payload
// @Description Information that should be available when you add a faculty
type AddFacultyRequestPayload struct {
	// Web Token that was appended to the link
	AddFacultyToken string

	// Faculty Name
	Name string `json:"name" validate:"required"`

	// Faculty Name Abbreviation
	Abbreviation string `json:"abbreviation" validate:"required"`
}
