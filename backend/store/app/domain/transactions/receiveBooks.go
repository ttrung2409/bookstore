package domain

import (
	module "store"
	"store/app/domain"
	data "store/data"
	"store/utils"
)

var transactionFactoryRef = module.Container.Get(utils.Nameof((*data.TransactionFactory)(nil))).(*data.TransactionFactory)
var transactionFactory = *transactionFactoryRef

func receiveBooks(books []data.Book) (*domain.BookReceipt, error) {
	transaction := transactionFactory.New()	

	var err error

	for _, book := range books {
		_, err = domain.CreateBookIfNotExist(book, transaction)
	}

	receipt, err := domain.CreateBookReceipt(books, transaction)

	if err != nil {
		(*transaction).Rollback()
		return nil, err
	}

	err = (*transaction).Commit()
	if (err != nil) {
		return nil, err
	}

 	return receipt, nil
}