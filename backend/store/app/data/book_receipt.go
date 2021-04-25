package data

import "time"

type BookReceipt struct {
	Id        EntityId
	StoreId   EntityId
	Number    string
	CreatedAt time.Time
}

type BookReceiptItem struct {
	BookReceiptId EntityId
	BookId        EntityId
	Qty           int
}

type BookReceiptRepository interface {
	repositoryBase
}

type BookReceiptItemRepository interface {
	repositoryBase
}
