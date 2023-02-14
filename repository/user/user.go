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

func (repo UserRepositoryImpl) Add(user user.User) error {
	return repo.db.Create(user).Error
}

func (repo UserRepositoryImpl) Get(username string) (*user.User, error) {
	result := &user.User{}
	err := repo.db.Where("username = ?", username).First(result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo UserRepositoryImpl) Update(user user.User) error {
	return repo.db.Save(user).Error
}

func (repo UserRepositoryImpl) Delete(username string) error {
	return repo.db.Where("username = ?", username).Delete(&user.User{}).Error
}
