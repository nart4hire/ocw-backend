package confirm

// Confirm Request Payload
// @Description Information that should be available when you confirm a password reset
type ConfirmRequestPayload struct {
	// Web Token that was appended to the link
	ConfirmToken string

	// User Password
	Password string `json:"password" validate:"required" example:"secret"`

	// User Password Validation, must be same as user
	PasswordValidation string `json:"password_validation" validate:"required,eqfield=Password" example:"secret"`
}
