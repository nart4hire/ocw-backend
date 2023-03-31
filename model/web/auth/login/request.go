package login

// Login Request Payload
//	@Description	Information that should be available when do a login process
type LoginRequestPayload struct {
	// User Email
	Email string `json:"email" validate:"required,email" example:"someone@example.com"`

	// User Password
	Password string `json:"password" validate:"required" example:"secret"`
}
