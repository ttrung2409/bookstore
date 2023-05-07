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
) (domain.Stock, error) {

	books, err := r.query(tx).
		Select("id", "onhand_qty", "reserved_qty").
		Where("id").In(bookIds).
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

func (r *bookRepository) adjustStock(
	adjustment domain.StockAdjustmentItem,
	tx repo.Transaction,
) error {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
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
