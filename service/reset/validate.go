package reset

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/validate"
)

func (rs ResetServiceImpl) Validate(payload validate.ValidateRequestPayload) error {
	// Double Layered Security
	// Validate Token
	_, err := rs.TokenUtil.Validate(payload.ValidateToken, token.Access)

	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Check if Token is Cached
	_, err = rs.CacheRepository.Get(*cache.NewKey(rs.RedisPrefixKey+"resetPassword", payload.ValidateToken))

	if err != nil {
		return web.NewResponseErrorFromError(err, web.LinkNotAvailable)
	}

	return nil
}
