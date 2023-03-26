package admin

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	req "gitlab.informatika.org/ocw/ocw-backend/model/web/admin/updateUser"
)

func (as AdminServiceImpl) UpdateUser(email string, payload req.AdminUpdateUserPayload) error {
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

	err := as.UserRepository.Update(user.User{
		Email:       payload.Email,
		Name:        payload.Name,
		Role:        role, // TODO: Change this
		IsActivated: false,
	})

	return err
}