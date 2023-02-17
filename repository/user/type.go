package user

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type UserRepository interface {
	Add(user user.User) error
	Get(username string) (*user.User, error)
	Update(user user.User) error
	Delete(username string) error
	IsExist(user string) (bool, error)
}
