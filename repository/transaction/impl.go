package transaction

import "gorm.io/gorm"

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db}
}

func (t *TransactionRepositoryImpl) Begin() {
	t.db = t.db.Begin()
}

func (t *TransactionRepositoryImpl) GetTransaction() *gorm.DB {
	return t.db
}

func (t *TransactionRepositoryImpl) Commit() {
	t.db.Commit()
}

func (t *TransactionRepositoryImpl) Rollback() {
	t.db.Rollback()
}
