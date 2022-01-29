package repository

import (
	"store/app/domain"
	"store/app/domain/data"
)

type BookReceiptRepository interface {
	repositoryBase
	Get(id data.EntityId, tx Transaction) (*domain.BookReceipt, error)
	Create(receipt *domain.BookReceipt, tx Transaction) (data.EntityId, error)
}
