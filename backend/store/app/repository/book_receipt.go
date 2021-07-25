package repository

import "store/app/domain/data"

type BookReceiptRepository interface {
	repositoryBase
	Get(id data.EntityId, tx Transaction) (*data.BookReceipt, error)
	Create(receipt *data.BookReceipt, tx Transaction) (data.EntityId, error)
}

type BookReceiptItemRepository interface {
	repositoryBase
}
