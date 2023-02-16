package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

type TokenUtilImpl struct {
	env.Environment
}

func (t TokenUtilImpl) DefaultMethod() jwt.SigningMethod {
	switch t.TokenMethod {
	case "hs256":
		return jwt.SigningMethodHS256
	default:
		return jwt.SigningMethodHS512
	}
}

func (tu TokenUtilImpl) Validate(tokenString string, tokenType token.TokenType) (*token.UserClaim, error) {
	jwtData, err := jwt.ParseWithClaims(tokenString, &token.UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		} else if method != tu.DefaultMethod() {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(tu.Environment.TokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtData.Claims.(*token.UserClaim)

	if !ok {
		return nil, fmt.Errorf("invalid claim")
	}

	if claims.Type != tokenType {
		return nil, fmt.Errorf("token type is not valid")
	}

	return claims, nil
}

func (t TokenUtilImpl) Generate(claim token.UserClaim, method jwt.SigningMethod) (string, error) {
	token := jwt.NewWithClaims(
		method,
		claim,
	)

	return token.SignedString([]byte(t.TokenSecret))
}
