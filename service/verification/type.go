package verification

type VerificationService interface {
	SendVerifyMail(email string) error
	SetVerification(email string, isVerified bool) error
}
