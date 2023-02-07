package user

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type UserRepository interface {
	Add(user user.User)
	Get(username string) user.User
	Update(user user.User)
	Delete(username string)
}
