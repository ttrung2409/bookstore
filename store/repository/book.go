package repository

import (
	"store/app/domain"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type BookRepository struct {
	postgresRepository[domain.BookData]
}

func (BookRepository) New() *BookRepository {
	return &BookRepository{postgresRepository: postgresRepository[domain.BookData]{db: GetDb()}}
}

func (r *BookRepository) CreateIfNotExist(
	book *domain.Book,
	tx *Transaction,
) error {
	bookData := book.State()

	db := r.db
	if tx != nil {
		db = tx.db
	}

	if result := db.Where("google_book_id = ?", bookData.GoogleBookId).FirstOrCreate(&bookData); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepository) GetStock(
	bookIds []string,
	tx *Transaction,
) (domain.Stock, error) {
	books, err := r.query(tx).
		Select("id", "onhand_qty", "reserved_qty").
		Where("id IN ?", bookIds).
		Find()

	if err != nil {
		return nil, err
	}

	stock := funk.Map(books, func(book domain.BookData) (string, domain.StockItem) {
		return book.Id, domain.StockItem{
			BookId:      book.Id,
			OnhandQty:   book.OnhandQty,
			ReservedQty: book.ReservedQty,
		}
	}).(domain.Stock)

	return stock, nil
}

func (r *BookRepository) adjustStock(
	adjustment domain.StockAdjustmentItem,
	tx *Transaction,
) error {
	db := GetDb()
	if tx != nil {
		db = tx.db
	}

	var result *gorm.DB

	switch adjustment.StockType {
	case domain.StockTypeOnhand:
		result = db.
			Model(&domain.BookData{}).
			Where("id = ?", adjustment.BookId).
			Update("onhand_qty", gorm.Expr("onhand_qty + ?", adjustment.Qty))
	case domain.StockTypeReserved:
		result = db.
			Model(&domain.BookData{}).
			Where("id = ?", adjustment.BookId).
			Update("reserved_qty", gorm.Expr("reserved_qty + ?", adjustment.Qty))
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
