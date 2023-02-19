package user

import (
	"errors"

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

func (repo UserRepositoryImpl) IsExist(email string) (bool, error) {
	_, err := repo.Get(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}

		return true, err
	}

	return false, nil
}

func (repo UserRepositoryImpl) Add(user user.User) error {
	return repo.db.Create(&user).Error
}

func (repo UserRepositoryImpl) Get(email string) (*user.User, error) {
	result := &user.User{}
	err := repo.db.Where("email = ?", email).First(result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo UserRepositoryImpl) GetAll() ([]user.User, error) {
	var result []user.User
	err := repo.db.Find(&result).Error

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
