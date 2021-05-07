package domain

import (
	module "store"
	data "store/app/data"
	"store/utils"
)

type Book struct {
	data.Book
}

var bookRepository = module.Container.Get(utils.Nameof((*data.BookRepository)(nil))).(data.BookRepository)

func (Book) CreateIfNotExist(book data.Book, transaction *data.Transaction) (data.BookId, error) {
	id, err := bookRepository.CreateIfNotExist(book, transaction)

	if err != nil {
		return data.EmptyBookId, err
	}

	return data.BookId(id.Value()), nil
}

func (book *Book) AdjustOnhandQty(qty int, transaction *data.Transaction) {
	bookRepository.AdjustOnhandQty(book.Id, qty, transaction)
}
