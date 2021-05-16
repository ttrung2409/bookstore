package domain

import (
	"errors"
	data "store/app/data"
)

type Book struct {
	data.Book
}

func (Book) Get(id data.EntityId, tx data.Transaction) (*Book, error) {
	record, err := BookRepository.Get(id, tx)
	if err != nil {
		return nil, err
	}

	return &Book{record.(data.Book)}, nil
}

func (Book) CreateIfNotExists(book data.Book, tx data.Transaction) (data.EntityId, error) {
	book.Id = data.NewEntityId()
	id, err := BookRepository.CreateIfNotExists(&book, tx)

	if err != nil {
		return data.EmptyEntityId, err
	}

	return id, nil
}

func (book *Book) AdjustOnhandQty(qty int, tx data.Transaction) error {
	if book.OnhandQty+qty < 0 {
		return errors.New("There is not enough stock")
	}

	return BookRepository.AdjustOnhandQty(book.Id, qty, tx)
}

func (book *Book) AdjustPreservedQty(qty int, tx data.Transaction) error {
	return BookRepository.AdjustPreservedQty(book.Id, qty, tx)
}
