package password

import (
	"fmt"

	"gitlab.informatika.org/ocw/ocw-backend/utils/base64"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"golang.org/x/crypto/bcrypt"
)

type PasswordUtilImpl struct {
	*env.Environment
	base64.Base64Util
}

func (e PasswordUtilImpl) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), e.Environment.PasswordCost)
	return string(hash), err
}

func (e PasswordUtilImpl) Check(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return fmt.Errorf("password mismatch")
	}

	return nil
}
