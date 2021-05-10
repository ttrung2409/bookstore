package postgres

import (
	data "store/app/data"

	"gorm.io/gorm"
)

type bookRepository struct {
	postgresRepository
}

func (r *bookRepository) CreateIfNotExist(
	book data.Book,
	tx *transaction,
) (data.EntityId, error) {
	db := Db()
	if tx.db != nil {
		db = tx.db
	}

	if result := db.Where("google_book_id = ?", book.GoogleBookId).FirstOrCreate(&book); result.Error != nil {
		return data.EmptyEntityId, result.Error
	}

	return book.Id, nil
}

func (r *bookRepository) AdjustOnhandQty(
	id data.EntityId,
	qty int,
	tx *transaction,
) error {
	db := Db()
	if tx != nil {
		db = tx.db
	}

	result := db.
		Model(&data.Book{}).
		Where("id = ?", id).
		Update("onhand_qty", gorm.Expr("onhand_qty + ?", qty))

	if result.Error != nil {
		return result.Error
	}

	return nil
}

var bookRepositoryInstance = bookRepository{postgresRepository{newEntity: func() data.Entity {
	return &data.Book{}
}}}
