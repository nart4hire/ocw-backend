package transaction

import "gorm.io/gorm"

type Transaction interface {
	Begin()
	GetTransaction() *gorm.DB
	Commit()
	Rollback()
	Auto(*bool)
}

type TransactionBuilder interface {
	Build() Transaction
}
