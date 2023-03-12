package transaction

import "gorm.io/gorm"

type Transaction interface {
	Begin()
	GetTransaction() *gorm.DB
	Commit()
	Rollback()
}

type TransactionBuilder interface {
	Build() Transaction
}
