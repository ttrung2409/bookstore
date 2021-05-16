package domain

import "store/app/data"

type TransactionalFunc func(tx data.Transaction) (interface{}, error)

func RunInTransaction(fn TransactionalFunc, ambientTx data.Transaction) (interface{}, error) {
	tx := ambientTx
	if tx == nil {
		tx = TransactionFactory.New()
	}

	var err error
	defer func() {
		if err != nil && ambientTx == nil {
			tx.Rollback()
		}
	}()

	result, err := fn(tx)

	if ambientTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
