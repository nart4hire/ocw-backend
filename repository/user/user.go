package user

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/utils/db"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func New(
	db db.Database,
) *UserRepositoryImpl {
	return &UserRepositoryImpl{db.Connect()}
}

func (repo UserRepositoryImpl) Add(user user.User) {
	repo.db.Create(user)
}

func (repo UserRepositoryImpl) Get(username string) user.User {
	result := user.User{}
	repo.db.First(&result, "username = ?", username)

	return result
}

func (repo UserRepositoryImpl) Update(user user.User) {
	repo.db.Save(user)
}

func (repo UserRepositoryImpl) Delete(username string) {
	repo.db.Delete(&user.User{}, "username = ?", username)
}
