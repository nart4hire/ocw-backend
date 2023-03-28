package token

import (
	"github.com/golang-jwt/jwt/v4"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Role       user.UserRole `json:"role"`
	Type       TokenType     `json:"type"`
	IsVerified bool          `json:"is_verified"`
}
