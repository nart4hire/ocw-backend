package register

// Register Request Payload
// @Description Information that should be available when do a registration process
type RegisterRequestPayload struct {
	// User Email
	Email string `json:"email" validate:"required,email" example:"someone@example.com"`

	// User Password
	Password string `json:"password" validate:"required" example:"secret"`

	// User Password Validation, must be same as user
	PasswordValidation string `json:"password_validation" validate:"required,eqfield=Password" example:"secret"`

	// User name
	Name string `json:"name" validate:"required" example:"someone"`
}
