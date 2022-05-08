package repository

import (
	"store/app/domain"
	"store/app/domain/data"
)

type BookRepository interface {
	CreateIfNotExist(book *domain.Book, tx Transaction) error
	GetStock(ids []string) data.Stock
}
