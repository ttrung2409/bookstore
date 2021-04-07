package domain

import (
	module "store"
	data "store/data"
	"store/utils"
)

type Book struct {
	data.Book
}

var bookRepositoryRef = module.Container.Get(utils.Nameof((*data.BookRepository)(nil))).(*data.BookRepository);
var bookRepository = *bookRepositoryRef;

func CreateBookIfNotExist(book data.Book, transaction *data.Transaction) (*Book, error) {
	err := bookRepository.CreateIfNotExist(book, transaction)

	if err != nil {
		return nil, err
	}

	createdBook := &Book{book}

	return createdBook, nil
}

func (book *Book) Update(entity data.Book) {
	bookRepository.Update(book, book, nil)
} 