package domain

import (
	module "store"
	"store/app/domain"
	data "store/data"
	"store/utils"
)

var transactionFactory = module.Container.Get(utils.Nameof((*data.TransactionFactory)(nil))).(data.TransactionFactory)

func receiveBooks(books []data.Book, qty map[string]int) (*domain.BookReceipt, error) {
	transaction := transactionFactory.New()

	var err error

	for _, book := range books {
		domainBook, err := domain.CreateBookIfNotExist(book, transaction)
		if err != nil {
			break
		}

		domainBook.UpdateOnhandQty(qty[domainBook.Id.Value().ToString()], transaction)
	}

	receipt, err := domain.CreateBookReceipt(books, qty, transaction)

	err = (*transaction).Commit()

	if err != nil {
		(*transaction).Rollback()

		return nil, err
	}

	return receipt, nil
}
