package auth

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/register"
)

func (auth AuthServiceImpl) Register(payload register.RegisterRequestPayload) error {
	hashedPassword, err := auth.Hash(payload.Password)

	if err != nil {
		return err
	}

	err = auth.UserRepository.Add(user.User{
		Email:       payload.Email,
		Password:    hashedPassword,
		Name:        payload.Name,
		Role:        user.Student,
		IsActivated: false,
	})

	if err == nil {
		auth.SendVerifyMail(payload.Email)
	}

	return err
}
