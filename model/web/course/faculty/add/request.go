package add

// AddFaculty Request Payload
// @Description Information that should be available when you add a faculty
type AddFacultyRequestPayload struct {
	// Faculty Name
	Name string `json:"name" validate:"required"`

	// Faculty Name Abbreviation
	Abbreviation string `json:"abbreviation" validate:"required"`
}
