package repository

import (
	repo "store/app/repository"

	"gorm.io/gorm"
)

type transaction struct {
	db *gorm.DB
}

type transactionFactory struct{}

func (f *transactionFactory) New() repo.Transaction {
	var tx *gorm.DB
	if tx := Db().Begin(); tx.Error != nil {
		return nil
	}

	return &transaction{tx}
}

func (f *transactionFactory) RunInTransaction(
	fn repo.TransactionalFunc,
) (interface{}, error) {
	tx := f.New()

	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	result, err := fn(tx)

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
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
