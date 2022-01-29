package repository

import (
	"store/app/domain"
	"store/app/domain/data"
)

type BookRepository interface {
	repositoryBase
	CreateIfNotExists(book *domain.Book, tx Transaction) (data.EntityId, error)
}
