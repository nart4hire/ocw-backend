package admin

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	addUser "gitlab.informatika.org/ocw/ocw-backend/model/web/admin/addUser"
	updateUser "gitlab.informatika.org/ocw/ocw-backend/model/web/admin/updateUser"
)

type AdminService interface {
	GetAllUser() ([]user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	AddUser(payload addUser.AdminAddUserPayload) error
	UpdateUser(email string, payload updateUser.AdminUpdateUserPayload) error
	DeleteUser(email string) error
}
