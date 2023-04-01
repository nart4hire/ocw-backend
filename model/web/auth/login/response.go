package login

// Login Response Payload
//	@Description	Login response when process success
type LoginResponsePayload struct {
	// Token that used to generate new access token
	RefreshToken string `json:"refresh_token"`

	// Token that used to access the resources
	AccessToken string `json:"access_token"`
}
