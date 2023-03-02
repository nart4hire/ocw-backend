package admin

// AdminUpdateUserPayload Request Payload
// @Description Information that should be available when admin update user

type AdminUpdateUserPayload struct {
	// User name
	Name string `json:"name" validate:"required" example:"someone"`

	// User Email
	Email string `json:"email" validate:"required,email" example:"someone@example.com"`

	// User Role
	Role string `json:"role" validate:"required" example:"admin"`
}
