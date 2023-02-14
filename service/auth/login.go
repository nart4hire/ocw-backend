package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/login"
	tokenModel "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gorm.io/gorm"
)

func (auth AuthServiceImpl) Login(payload login.LoginRequestPayload) (*login.LoginResponsePayload, error) {
	user, err := auth.Get(payload.Email)

	if err != nil {
		var errorObj error

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			errorObj = fmt.Errorf("username or password combination not found")
		default:
			errorObj = err
		}

		return nil, errorObj
	}

	if err := auth.Check(payload.Password, user.Password); err != nil {
		return nil, err
	}

	refreshClaim := tokenModel.UserClaim{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Type:  tokenModel.Refresh,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(auth.TokenRefreshExpired) * time.Millisecond)),
			Issuer:    auth.TokenIssuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessClaim := tokenModel.UserClaim{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Type:  tokenModel.Access,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(auth.TokenAccessExpired) * time.Millisecond)),
			Issuer:    auth.TokenIssuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken, err := auth.TokenUtil.Generate(refreshClaim)

	if err != nil {
		return nil, err
	}

	accessToken, err := auth.TokenUtil.Generate(accessClaim)

	if err != nil {
		return nil, err
	}

	return &login.LoginResponsePayload{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
