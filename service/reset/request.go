package reset

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	tokenModel "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/provider/mail"
	"gorm.io/gorm"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/request"
)

func (rs ResetServiceImpl) Request(payload request.RequestRequestPayload) error {
	// Fetch user data from email
	user, err := rs.UserRepository.Get(payload.Email)

	if err != nil {
		var errorObj error

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			errorObj = web.NewResponseError("Email was not found", web.EmailNotExist)
		default:
			errorObj = err
		}

		return errorObj
	}

	if !user.IsActivated {
		return web.NewResponseError("user is not activated yet", web.InactiveUser)
	}

	// Mint JWT Token for 30 minutes
	resetClaim := tokenModel.UserClaim{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Type:  tokenModel.Access,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(rs.TokenAccessExpired*6) * time.Millisecond)),
			Issuer:    rs.TokenIssuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	resetToken, err := rs.TokenUtil.Generate(resetClaim, rs.TokenUtil.DefaultMethod())
	if err != nil {
		return err
	}

	// Cache Website on Redis, TTL 30 mins
	rs.CacheRepository.Set(*cache.NewString(*cache.NewKey(rs.RedisPrefixKey+"resetPassword", resetToken), payload.Email, 30))

	// Send Reset Email
	mailBuilder, err := rs.TemplateWritterBuilder.Get("reset-password.format.html")

	if err != nil {
		return err
	}

	mailData, err := mailBuilder.Write(&mailPayload{
		BaseUrl: rs.FrontendBaseURL + rs.ResetPasswordPath,
		Email:   user.Email,
		Token:   resetToken,
	})

	if err != nil {
		return err
	}

	rs.MailQueue.Send(mail.Mail{
		To:      []string{user.Email},
		Subject: "Reset Password",
		Message: mailData,
	})

	return nil
}
