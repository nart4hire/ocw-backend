package request

// Request Request Payload
// @Description Information that should be available when password reset is requested
type RequestRequestPayload struct {
	// User Email
	Email string `json:"email" validate:"required,email" example:"someone@example.com"`
}
