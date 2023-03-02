package verification

type VerificationService interface {
	SendVerifyMail(email string) error
	DoVerification(id string) error
}
