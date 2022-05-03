package repository

import (
	"store/app/domain"
)

type BookReceiptRepository interface {
	Get(id string, tx Transaction) (*domain.BookReceipt, error)
	Create(receipt *domain.BookReceipt, tx Transaction) (string, error)
}
