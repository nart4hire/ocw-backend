package token_test

import (
	"testing"

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

		token, err := tokenObj.Generate(claim)
		assert.Nil(t, err)

		extractedToken, err := tokenObj.Validate(token, tokenData.Refresh)
		assert.Nil(t, err)
		assert.NotNil(t, extractedToken)

		assert.Equal(t, claim, *extractedToken)
	})

}
