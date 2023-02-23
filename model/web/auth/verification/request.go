package verification

// Email Verification Request Payload
// @Description Information that should be passed when request verify
type VerificationRequestPayload struct {
	// User Email
	Email string `json:"email" validate:"required,email" example:"someone@example.com"`
}
