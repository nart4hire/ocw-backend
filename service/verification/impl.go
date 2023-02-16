package verification

import "gitlab.informatika.org/ocw/ocw-backend/provider/mail"

type VerificationServiceImpl struct {
	mail.MailQueue
}
