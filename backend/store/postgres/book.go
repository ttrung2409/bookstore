package postgres

import (
	data "store/app/data"
)

type bookRepository struct {
	postgresRepository
}

func (r *bookRepository) CreateIfNotExist(
	book *data.Book,
	tx data.Transaction,
) (data.EntityId, error) {
	book.Id = data.NewEntityId()

	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	if result := db.Where("google_book_id = ?", book.GoogleBookId).FirstOrCreate(book); result.Error != nil {
		return data.EmptyEntityId, result.Error
	}

	return book.Id, nil
}

var bookRepositoryInstance = bookRepository{postgresRepository{newEntity: func() data.Entity {
	return &data.Book{}
}}}
