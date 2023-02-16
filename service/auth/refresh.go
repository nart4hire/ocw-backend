package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/refresh"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
)

func (auth AuthServiceImpl) Refresh(payload refresh.RefreshRequestPayload) (*refresh.RefreshResponsePayload, error) {
	claim, err := auth.TokenUtil.Validate(payload.RefreshToken, token.Refresh)

	if err != nil {
		return nil, err
	}

	claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(auth.TokenAccessExpired) * time.Millisecond))
	claim.Type = token.Access

	newToken, err := auth.TokenUtil.Generate(*claim)

	if err != nil {
		return nil, err
	}

	return &refresh.RefreshResponsePayload{
		AccessToken: newToken,
	}, nil
}
