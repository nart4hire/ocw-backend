package verification

// Email Verification Request Payload
// @Description Information that should be passed when request verify
type VerificationSendRequestPayload struct {
	// User Email
	Email string `json:"email" validate:"required,email" example:"someone@example.com"`
}

type VerificationRequestPayload struct {
	Id string `json:"id" validate:"required" example:"6ba7b812-9dad-11d1-80b4-00c04fd430c8"`
}
