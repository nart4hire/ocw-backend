package refresh

// Refresh Response Payload
//	@Description	Refresh endpoint response when process success
type RefreshResponsePayload struct {
	// Token that used to access the resources
	AccessToken string `json:"access_token"`
}
