package password_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.informatika.org/ocw/ocw-backend/utils/password"
)

func TestPasswordHash(t *testing.T) {
	obj := password.PasswordUtilImpl{}

	t.Run("PasswordCanBeHashed", func(t *testing.T) {
		_, err := obj.Hash("admin")

		assert.Nil(t, err)
	})

	t.Run("PasswordHashMustBeDifferOnSamePass", func(t *testing.T) {
		hash1, err := obj.Hash("admin")
		assert.Nil(t, err)

		hash2, err := obj.Hash("admin")
		assert.Nil(t, err)

		assert.NotEqual(t, hash1, hash2)
	})

	t.Run("PasswordCanBeHashAndValidateCorrectly", func(t *testing.T) {
		hash, err := obj.Hash("admin")
		assert.Nil(t, err)

		err = obj.Check("admin", hash)
		assert.Nil(t, err)

		err = obj.Check("seseorang", hash)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "password mismatch")
	})
}
