package token_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	tokenData "gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
)

func TestToken(t *testing.T) {
	tokenObj := token.TokenUtilImpl{
		Environment: env.Environment{
			TokenSecret: "secret",
			TokenMethod: "hs512",
		},
	}

	t.Run("UserTokenTest", func(t *testing.T) {
		claim := tokenData.UserClaim{
			Name:  "Someone",
			Email: "someone@example.com",
			Role:  user.Member,
			Type:  tokenData.Refresh,
		}

		token, err := tokenObj.Generate(claim, tokenObj.DefaultMethod())
		assert.Nil(t, err)

		extractedToken, err := tokenObj.Validate(token, tokenData.Refresh)
		assert.Nil(t, err)
		assert.NotNil(t, extractedToken)

		assert.Equal(t, claim, *extractedToken)
	})

	t.Run("UserTokenInvalidType", func(t *testing.T) {
		claim := tokenData.UserClaim{
			Name:  "Someone",
			Email: "someone@example.com",
			Role:  user.Member,
			Type:  tokenData.Refresh,
		}

		token, err := tokenObj.Generate(claim, tokenObj.DefaultMethod())
		assert.Nil(t, err)

		extractedToken, err := tokenObj.Validate(token, tokenData.Access)
		assert.NotNil(t, err)
		assert.Nil(t, extractedToken)
		assert.Equal(t, err.Error(), "token type is not valid")
	})

	t.Run("UserTokenExpired", func(t *testing.T) {
		claim := tokenData.UserClaim{
			Name:  "Someone",
			Email: "someone@example.com",
			Role:  user.Member,
			Type:  tokenData.Refresh,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now()),
			},
		}

		token, err := tokenObj.Generate(claim, tokenObj.DefaultMethod())
		assert.Nil(t, err)

		extractedToken, err := tokenObj.Validate(token, tokenData.Refresh)
		assert.NotNil(t, err)
		assert.Nil(t, extractedToken)
		assert.Contains(t, err.Error(), "expired")
	})
}
