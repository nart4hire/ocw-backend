package auth

import "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/verification"

func (auth AuthServiceImpl) SendVerifyEmail(payload verification.VerificationSendRequestPayload) error {
	return auth.VerificationService.SendVerifyMail(payload.Email)
}
