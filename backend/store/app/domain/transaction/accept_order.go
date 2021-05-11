package domain

import (
	"store/app/data"
	"store/app/domain"
)

func AcceptOrder(id data.EntityId) error {
	tx := domain.TransactionFactory.New()
	var err error

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	order, err := domain.Order{}.Get(id, tx)
	if err != nil {
		return err
	}

	if err = order.Accept(tx); err != nil {
		return err
	}

	return nil
}
