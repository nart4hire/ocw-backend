package admin

import (
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
)

type AdminServiceImpl struct {
	UserRepository user.UserRepository
}