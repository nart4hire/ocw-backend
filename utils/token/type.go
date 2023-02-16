package token

import (
	"github.com/golang-jwt/jwt/v4"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
)

type TokenUtil interface {
	Validate(jwt string, tokenType token.TokenType) (*token.UserClaim, error)
	Generate(claim token.UserClaim, method jwt.SigningMethod) (string, error)
	DefaultMethod() jwt.SigningMethod
}
