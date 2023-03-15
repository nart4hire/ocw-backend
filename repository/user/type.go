package user

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
)

type UserRepository interface {
	Add(user user.User) error
	Get(username string) (*user.User, error)
	GetAll() ([]user.User, error)
	Update(user user.User) error
	Delete(username string) error
	IsExist(user string) (bool, error)

	AddWithTransaction(tx transaction.Transaction, user user.User) error
	GetWithTransaction(tx transaction.Transaction, username string) (*user.User, error)
	GetAllWithTransaction(tx transaction.Transaction) ([]user.User, error)
	UpdateWithTransaction(tx transaction.Transaction, user user.User) error
	DeleteWithTransaction(tx transaction.Transaction, username string) error
	IsExistWithTransaction(tx transaction.Transaction, user string) (bool, error)
}
