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
) (string, bool, error) {
	dataBook := book.State()

	newEntityId := data.NewEntityId()
	dataBook.Id = newEntityId

	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	if result := db.Where("google_book_id = ?", dataBook.GoogleBookId).FirstOrCreate(dataBook); result.Error != nil {
		return data.EmptyEntityId, false, result.Error
	}

	return dataBook.Id, dataBook.Id == newEntityId, nil
}

func (r *bookRepository) adjustOnhandQty(
	id string,
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
