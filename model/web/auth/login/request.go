package login

type LoginRequestPayload struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	CaptchaToken string `json:"token"`
}
