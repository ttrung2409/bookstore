package repository

import (
	"store/app/domain"
	repo "store/app/repository"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type bookRepository struct {
	postgresRepository[domain.BookData]
}

func (r *bookRepository) CreateIfNotExist(
	book *domain.Book,
	tx repo.Transaction,
) error {
	bookData := book.State()

	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	if result := db.Where("google_book_id = ?", bookData.GoogleBookId).FirstOrCreate(&bookData); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *bookRepository) GetStock(
	bookIds string,
	tx repo.Transaction,
) (*domain.Stock, error) {

	books, err := r.query(tx).
		Select("id", "onhand_qty", "reserved_qty").
		Where("id").In(bookIds).
		Find()

	if err != nil {
		return nil, err
	}

	stock := funk.Map(books, func(book domain.BookData) (string, domain.StockItemData) {
		return book.Id, domain.StockItemData{
			BookId:      book.Id,
			OnhandQty:   book.OnhandQty,
			ReservedQty: book.ReservedQty,
		}
	}).(domain.StockData)

	return domain.Stock{}.New(stock), nil
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
		Model(&domain.BookData{}).
		Where("id = ?", id).
		Update("onhand_qty", gorm.Expr("onhand_qty + ?", qty))

	if result.Error != nil {
		return result.Error
	}

	return nil
}
