package password

import (
	"fmt"

	"gitlab.informatika.org/ocw/ocw-backend/utils/base64"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"golang.org/x/crypto/bcrypt"
)

type PasswordUtilImpl struct {
	env.Environment
	base64.Base64Util
}

func (e PasswordUtilImpl) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), e.Environment.PasswordCost)
	return e.Base64Util.Encode(hash), err
}

func (e PasswordUtilImpl) Check(password string, hashedPassword string) error {
	hash, err := e.Base64Util.Decode(hashedPassword)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))

	if err != nil {
		return fmt.Errorf("username or password combination is not found")
	}

	return nil
}
