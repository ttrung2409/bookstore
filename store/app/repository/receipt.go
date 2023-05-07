package repository

import (
	"store/app/domain"
)

type ReceiptRepository interface {
	Create(receipt *domain.Receipt, tx Transaction) error
}
