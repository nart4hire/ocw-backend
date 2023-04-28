package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail/smtp"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

func DisabledTestSendMail(t *testing.T) {
	smtpClient := smtp.New(&env.Environment{
		SmtpUsername: "",
		SmtpPassword: "",
		SmtpIdentity: "",
		SmtpServer:   "",
		SmtpPort:     21,
		SmtpAuthType: "plain",
	})

	err := smtpClient.Send([]string{"bayusamudra.55.02.com@gmail.com"}, "Testing", "Ini test")

	assert.Nil(t, err)
}
