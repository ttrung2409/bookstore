package domain

import (
	module "store"
	data "store/data"
	"store/utils"
)

type Book struct {
	data.Book
}

var bookRepository = module.Container.Get(utils.Nameof((*data.BookRepository)(nil))).(data.BookRepository)

func CreateBookIfNotExist(book data.Book, transaction *data.Transaction) (*Book, error) {
	id, err := bookRepository.CreateIfNotExist(book, transaction)

	if err != nil {
		return nil, err
	}

	createdBook := &Book{book}
	createdBook.Id = data.BookId(id.Value())

	return createdBook, nil
}

func (book *Book) UpdateOnhandQty(qty int, transaction *data.Transaction) {
	bookRepository.Update(book.Id, qty, transaction)
}
