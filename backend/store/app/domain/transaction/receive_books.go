package domain

import (
	module "store"
	data "store/app/data"
	"store/app/domain"
	"store/utils"
)

var transactionFactory = module.Container.Get(utils.Nameof((*data.TransactionFactory)(nil))).(data.TransactionFactory)

func ReceiveBooks(items []domain.ReceivingBook) (*domain.BookReceipt, error) {
	transaction := transactionFactory.New()

	var err error

	for _, item := range items {
		_, err := domain.Book{}.CreateIfNotExist(item.Book, transaction)
		if err != nil {
			break
		}

		book.AdjustOnhandQty(item.Qty, transaction)
	}

	receipt, err := domain.BookReceipt{}.Create(items, transaction)

	err = (*transaction).Commit()

	if err != nil {
		(*transaction).Rollback()

		return nil, err
	}

	return receipt, nil
}
