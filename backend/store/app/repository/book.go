package repository

import "store/app/domain/data"

type BookRepository interface {
	repositoryBase
	CreateIfNotExists(book *data.Book, tx Transaction) (data.EntityId, error)
}
