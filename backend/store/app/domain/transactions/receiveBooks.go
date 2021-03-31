package domain

import (
	module "store"
	"store/app/domain"
	repository "store/repository/interface"
	"store/utils"
)

var transactionFactoryRef = module.Container.Get(utils.Nameof((*repository.TransactionFactory)(nil))).(*repository.TransactionFactory)
var transactionFactory = *transactionFactoryRef

func receiveBooks(books []domain.Book) {
	transaction := transactionFactory.New()

	for _, book := range books {
		domain.CreateBookIfNotExist(book, transaction)
	}

	domain.CreateBookReceipt(books, transaction)
	
	transaction.Commit()
}