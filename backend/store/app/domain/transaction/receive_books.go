package domain

import (
	module "store"
	data "store/app/data"
	"store/app/domain"
	"store/utils"
)

var transactionFactory = module.Container().Get(utils.Nameof((*data.TransactionFactory)(nil))).(data.TransactionFactory)

func ReceiveBooks(books []domain.ReceivingBook) (*domain.BookReceipt, error) {
	tx := transactionFactory.New()
	var err error

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, receivingBook := range books {
		book, err := domain.Book{}.CreateIfNotExists(receivingBook.Book, tx)
		if err != nil {
			return nil, err
		}

		err = book.AdjustOnhandQty(receivingBook.ReceivingQty, tx)
		if err != nil {
			return nil, err
		}
	}

	receipt, err := domain.BookReceipt{}.Create(books, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return receipt, nil
}
