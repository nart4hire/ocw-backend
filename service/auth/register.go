package auth

import (
	"errors"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/register"
	"gorm.io/gorm"
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

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err := auth.SendVerifyMail(payload.Email)
		if err != nil {
			auth.Logger.Warning("Failed to send email: " + err.Error())
		}
		return nil
	}

	if err == nil {
		err := auth.SendVerifyMail(payload.Email)
		if err != nil {
			auth.Logger.Warning("Failed to send email: " + err.Error())
		}
	}

	return err
}
