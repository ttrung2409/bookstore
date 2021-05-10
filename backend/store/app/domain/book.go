package domain

import (
	module "store"
	data "store/app/data"
	"store/utils"
)

type Book struct {
	data.Book
}

var bookRepository = module.Container().Get(utils.Nameof((*data.BookRepository)(nil))).(data.BookRepository)

func (Book) CreateIfNotExists(book data.Book, tx data.Transaction) (*Book, error) {
	book.Id = data.NewEntityId()
	_, err := bookRepository.CreateIfNotExists(book, tx)

	if err != nil {
		return nil, err
	}

	return &Book{book}, nil
}

func (book *Book) AdjustOnhandQty(qty int, tx data.Transaction) error {
	return bookRepository.AdjustOnhandQty(book.Id, qty, tx)
}
