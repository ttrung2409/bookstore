package repository

import (
	"store/app/domain"
)

type BookRepository interface {
	CreateIfNotExist(book *domain.Book, tx Transaction) error
	GetStock(bookIds []string) domain.Stock
}
