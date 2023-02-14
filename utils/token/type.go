package token

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
)

type TokenUtil interface {
	Validate(jwt string, tokenType token.TokenType) (*token.UserClaim, error)
	Generate(claim token.UserClaim) (string, error)
}
