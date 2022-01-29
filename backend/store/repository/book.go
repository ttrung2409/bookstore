package repository

import (
	"store/app/domain"
	data "store/app/domain/data"
	repo "store/app/repository"

	"gorm.io/gorm"
)

type bookRepository struct {
	postgresRepository
}

func (r *bookRepository) CreateIfNotExist(
	book *domain.Book,
	tx repo.Transaction,
) (data.EntityId, error) {
	dataBook := book.State()

	dataBook.Id = data.NewEntityId()

	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	if result := db.Where("google_book_id = ?", dataBook.GoogleBookId).FirstOrCreate(dataBook); result.Error != nil {
		return data.EmptyEntityId, result.Error
	}

	return dataBook.Id, nil
}

func (r *bookRepository) AdjustOnhandQty(
	id data.EntityId,
	qty int,
	tx repo.Transaction,
) error {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
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

func (r *bookRepository) AdjustReservedQty(
	id data.EntityId,
	qty int,
	tx repo.Transaction,
) error {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	result := db.
		Model(&data.Book{}).
		Where("id = ?", id).
		Update("reserved_qty", gorm.Expr("reservedQty + ?", qty))

	if result.Error != nil {
		return result.Error
	}

	return nil
}

var bookRepositoryInstance = bookRepository{postgresRepository{newEntity: func() data.Entity {
	return &data.Book{}
}}}
