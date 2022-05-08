package repository

import (
	"store/app/domain"
	data "store/app/domain/data"
	repo "store/app/repository"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type bookRepository struct {
	postgresRepository[data.Book]
}

func (r *bookRepository) CreateIfNotExist(
	book *domain.Book,
	tx repo.Transaction,
) error {
	dataBook := book.State()

	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	if result := db.Where("google_book_id = ?", dataBook.GoogleBookId).FirstOrCreate(&dataBook); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *bookRepository) GetStock(
	ids string,
	tx repo.Transaction,
) (data.Stock, error) {

	books, err := r.query(tx).
		Select("id", "onhand_qty", "reserved_qty").
		Where("id").In(ids).
		Find()

	if err != nil {
		return nil, err
	}

	return funk.Map(books, func(book data.Book) (string, data.StockItem) {
		return book.Id, data.StockItem{
			BookId:      book.Id,
			OnhandQty:   book.OnhandQty,
			ReservedQty: book.ReservedQty,
		}
	}).(data.Stock), nil
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
