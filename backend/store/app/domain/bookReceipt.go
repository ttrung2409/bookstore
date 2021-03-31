package domain

import (
	repository "store/repository/interface"
	"time"
)

type BookReceipt struct {
	Id repository.EntityId
	Number string
	CreatedAt time.Time
}

type BookReceiptItem struct {
	BookReceiptId repository.EntityId
	BookId repository.EntityId
	Qty int
}

func CreateBookReceipt(books []Book, transaction repository.Transaction) {

}