package admin

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type AdminService interface {
	GetAllUser() ([]user.User, error)
	GetUserByEmail() string
	AddUser() string
	UpdateUser() string
	DeleteUser() string
}
