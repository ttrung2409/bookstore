package domain

import (
	data "store/app/data"
)

type BookReceipt struct {
	state *data.BookReceipt
}

type ReceivingBook struct {
	data.Book
	ReceivingQty int
}

func (BookReceipt) New(state *data.BookReceipt) *BookReceipt {
	receipt := &BookReceipt{state: state}
	return receipt
}

func (BookReceipt) NewFromReceivingBooks(books []*ReceivingBook) *BookReceipt {
	receipt := &data.BookReceipt{
		Id: data.NewEntityId(),
	}

	items := []data.BookReceiptItem{}

	for _, book := range books {
		item := data.BookReceiptItem{
			Id:            data.NewEntityId(),
			BookReceiptId: receipt.Id,
			BookId:        book.Id,
			Qty:           book.ReceivingQty,
		}

		items = append(items, item)
	}

	receipt.Items = items

	return &BookReceipt{state: receipt}
}

func (receipt *BookReceipt) State() *data.BookReceipt {
	return receipt.state.Clone()
}
