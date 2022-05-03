package domain

import data "store/app/domain/data"

type Book struct {
	state data.Book
}

func (Book) New(state data.Book) *Book {
	return &Book{state: state}
}

func (b *Book) State() data.Book {
	return b.state.Clone()
}
