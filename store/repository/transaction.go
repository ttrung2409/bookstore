package repository

import (
	"gorm.io/gorm"
)

type Transaction struct {
	db *gorm.DB
}

type TransactionalFunc func(tx *Transaction) (any, error)

func (Transaction) New() *Transaction {
	var tx *gorm.DB
	if tx := GetDb().Begin(); tx.Error != nil {
		return nil
	}

	return &Transaction{tx}
}

func (Transaction) RunInTransaction(
	fn TransactionalFunc,
) (any, error) {
	tx := Transaction{}.New()

	var err error
	defer func() {
		if err != nil {
			tx.rollback()
		}
	}()

	result, err := fn(tx)

	if err != nil {
		return nil, err
	}

	err = tx.commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (tx *Transaction) commit() error {
	if result := tx.db.Commit(); result.Error != nil {
		return result.Error
	}

	return nil
}

func (tx *Transaction) rollback() error {
	if result := tx.db.Rollback(); result.Error != nil {
		return result.Error
	}

	return nil
}
