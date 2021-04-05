package data

import "time"

type BookReceipt struct {
	Id        EntityId
	Number    string
	CreatedAt time.Time
}

type BookReceiptItem struct {
	BookReceiptId EntityId
	BookId        EntityId
	Qty           int
}
