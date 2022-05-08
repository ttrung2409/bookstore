package domain

import data "store/app/domain/data"

type Book struct {
	state data.Book
}

func (Book) New(state data.Book) *Book {
	cloned := state.Clone()
	if cloned.Id == "" {
		cloned.Id = data.NewId()
	}

	return &Book{state: state.Clone()}
}

func (b *Book) State() data.Book {
	return b.state.Clone()
}
