package repository

import (
	"store/app/domain"
)

type ReceiptRepository interface {
	Get(id string, tx Transaction) (*domain.Receipt, error)
	Create(receipt *domain.Receipt, tx Transaction) error
}
