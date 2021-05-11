package postgres

import (
	"store/app/data"

	"gorm.io/gorm"
)

type transaction struct {
	db *gorm.DB
}

type transactionFactory struct{}

func (f *transactionFactory) New() data.Transaction {
	var tx *gorm.DB
	if tx := Db().Begin(); tx.Error != nil {
		return nil
	}

	return &transaction{tx}
}

func (tx *transaction) Commit() error {
	if result := tx.db.Commit(); result.Error != nil {
		return result.Error
	}

	return nil
}

func (tx *transaction) Rollback() error {
	if result := tx.db.Rollback(); result.Error != nil {
		return result.Error
	}

	return nil
}
