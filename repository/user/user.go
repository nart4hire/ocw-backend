package user

import (
	"errors"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	builder transaction.TransactionBuilder
}

func New(
	builder transaction.TransactionBuilder,
) *UserRepositoryImpl {
	return &UserRepositoryImpl{builder}
}

func (repo UserRepositoryImpl) IsExist(email string) (bool, error) {
	return repo.IsExistWithTransaction(repo.builder.Build(), email)
}

func (repo UserRepositoryImpl) IsExistWithTransaction(tx transaction.Transaction, email string) (bool, error) {
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
	return repo.AddWithTransaction(repo.builder.Build(), user)
}

func (repo UserRepositoryImpl) AddWithTransaction(tx transaction.Transaction, user user.User) error {
	return tx.GetTransaction().Create(&user).Error
}

func (repo UserRepositoryImpl) Get(email string) (*user.User, error) {
	return repo.GetWithTransaction(repo.builder.Build(), email)
}

func (repo UserRepositoryImpl) GetWithTransaction(tx transaction.Transaction, email string) (*user.User, error) {
	result := &user.User{}
	err := tx.GetTransaction().Where("email = ?", email).First(result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo UserRepositoryImpl) GetAll() ([]user.User, error) {
	return repo.GetAllWithTransaction(repo.builder.Build())
}
func (repo UserRepositoryImpl) GetAllWithTransaction(tx transaction.Transaction) ([]user.User, error) {
	var result []user.User
	err := tx.GetTransaction().Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo UserRepositoryImpl) Update(user user.User) error {
	return repo.UpdateWithTransaction(repo.builder.Build(), user)
}

func (repo UserRepositoryImpl) UpdateWithTransaction(tx transaction.Transaction, user user.User) error {
	return tx.GetTransaction().Save(user).Error
}

func (repo UserRepositoryImpl) Delete(email string) error {
	return repo.DeleteWithTransaction(repo.builder.Build(), email)
}

func (repo UserRepositoryImpl) DeleteWithTransaction(tx transaction.Transaction, email string) error {
	return tx.GetTransaction().Where("email = ?", email).Delete(&user.User{}).Error
}
