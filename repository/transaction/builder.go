package transaction

import (
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
	"gorm.io/gorm"
)

type TransactionBuilderImpl struct {
	db *gorm.DB
}

func NewBuilder(
	db db.Database,
) *TransactionBuilderImpl {
	return &TransactionBuilderImpl{db.Connect()}
}

func (t TransactionBuilderImpl) Build() Transaction {
	return New(t.db)
}
