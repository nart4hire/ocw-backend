package reset

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/confirm"
)

func (rs ResetServiceImpl) Confirm(payload confirm.ConfirmRequestPayload) error {
	// Double Layered Security
	// Validate Token
	_, err := rs.TokenUtil.Validate(payload.ConfirmToken, token.Access)

	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Check if Token is Cached
	email, err := rs.CacheRepository.Get(*cache.NewKey(rs.RedisPrefixKey+"resetPassword", payload.ConfirmToken))

	if err != nil {
		return web.NewResponseErrorFromError(err, web.LinkNotAvailable)
	}

	// Reset the password
	user, err := rs.UserRepository.Get(email)

	if err != nil {
		return err
	}

	hashedPassword, err := rs.Hash(payload.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword
	err = rs.UserRepository.Update(*user)

	if err != nil {
		return err
	}

	return nil
}
