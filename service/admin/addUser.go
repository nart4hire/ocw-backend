package admin

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	req "gitlab.informatika.org/ocw/ocw-backend/model/web/admin/addUser"
)

func (as AdminServiceImpl) AddUser(payload req.AdminAddUserPayload) error {
	// change role payload from string to user.UserRole
	var role user.UserRole

	// TODO: move this
	if (payload.Role == "admin") {
		role = user.Admin
	} else if (payload.Role == "contributor") {
		role = user.Contributor
	} else if (payload.Role == "member") {
		role = user.Student
	}

	err := as.UserRepository.Add(user.User{
		Email:       payload.Email,
		Name:        payload.Name,
		Role:        role,
		IsActivated: false,
	})

	return err
}