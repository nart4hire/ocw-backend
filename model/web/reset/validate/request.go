package validate

// Validate Request Payload
// @Description Information that should be available when link validation is done
type ValidateRequestPayload struct {
	// Web Token that was appended to the link
	ValidateToken string
}
