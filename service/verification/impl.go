package verification

import (
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail"
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
)

type VerificationServiceImpl struct {
	mail.MailQueue
	user.UserRepository
}
