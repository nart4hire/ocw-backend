package verification

import (
	"time"

	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail"
)

func (v VerificationServiceImpl) SendVerifyMail(email string) error {
	res, err := v.CacheRepository.GetInteger(cache.Key{
		Id: v.RedisPrefixKey + "verify:cnt:" + email,
	})

	if err != nil || res > v.Environment.EmailVerificationMaxRetry {
		return err
	}

	id := uuid.New().String()
	v.CacheRepository.Incr(v.RedisPrefixKey+"verify:cnt"+email, v.EmailVerificationRetryInterval*int64(time.Minute))
	v.CacheRepository.Set(cache.String{
		Key: cache.Key{
			Id: v.RedisPrefixKey + "verify:id:" + id,
		},
		Value:           email,
		ExpiryInMinutes: int(v.EmailVerificationExpire) * int(time.Second),
	})

	mailBuilder, err := v.TemplateWritterBuilder.Get("email-verification.format.html")

	if err != nil {
		return err
	}

	mailData, err := mailBuilder.Write(&mailPayload{
		BaseUrl: v.FrontendBaseURL + v.ResetPasswordPath,
		Email:   email,
		Token:   id,
	})

	if err != nil {
		return err
	}

	v.MailQueue.Send(mail.Mail{
		To:      []string{email},
		Subject: "Email Verification",
		Message: mailData,
	})

	return nil
}
