package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail/smtp"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

func TestSendMail(t *testing.T) {
	smtpClient := smtp.New(&env.Environment{
		SmtpUsername: "postmaster@bayusamudra.my.id",
		SmtpPassword: "4868b913c0a48289beb3eccd13a35de4-70c38fed-8afa51c1",
		SmtpIdentity: "",
		SmtpServer:   "smtp.mailgun.org",
		SmtpPort:     587,
		SmtpAuthType: "plain",
	})

	err := smtpClient.Send([]string{"bayusamudra.55.02.com@gmail.com"}, "Testing", "Ini test")

	assert.Nil(t, err)
}
