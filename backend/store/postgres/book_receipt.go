package postgres

import "store/app/data"

type bookReceiptRepository struct {
	postgresRepository
}

type bookReceiptItemRepository struct {
	postgresRepository
}

var bookReceiptRepositoryInstance = bookRepository{postgresRepository{newEntity: func() interface{} {
	return &data.BookReceipt{}
}}}

var bookReceiptItemRepositoryInstance = bookRepository{postgresRepository{newEntity: func() interface{} {
	return &data.BookReceiptItem{}
}}}
