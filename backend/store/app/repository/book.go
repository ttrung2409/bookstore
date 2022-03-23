package repository

import (
	"store/app/domain"
)

type BookRepository interface {
	repositoryBase
	CreateIfNotExists(book *domain.Book, tx Transaction) (string, bool, error)
}
