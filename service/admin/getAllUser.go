package admin

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

func (as AdminServiceImpl) GetAllUser() ([]user.User, error) {
	var users []user.User
	users, nil := as.UserRepository.GetAll()
	return users, nil
}