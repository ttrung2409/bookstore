package domain

import (
	data "store/app/domain/data"

	"github.com/thoas/go-funk"
)

type BookReceipt struct {
	state data.BookReceipt
}

type ReceivingBook struct {
	data.Book
	ReceivingQty int
}

func (BookReceipt) New(state data.BookReceipt) *BookReceipt {
	cloned := state.Clone()
	if cloned.Id == "" {
		cloned.Id = data.NewId()
	}

	receipt := &BookReceipt{state: cloned}
	return receipt
}

func (BookReceipt) NewFromReceivingBooks(books []ReceivingBook) *BookReceipt {
	receipt := data.BookReceipt{
		Id: data.NewId(),
	}

	items := []data.BookReceiptItem{}

	for _, book := range books {
		items = append(items, data.BookReceiptItem{
			BookReceiptId: receipt.Id,
			BookId:        book.Id,
			Qty:           book.ReceivingQty,
		})
	}

	receipt.Items = items
	receipt.OnhandStockAdjustment = funk.Map(receipt.Items, func(item data.BookReceiptItem) data.StockAdjustmentItem {
		return data.StockAdjustmentItem{BookId: item.BookId, Qty: item.Qty}
	}).(data.StockAdjustment)

	return &BookReceipt{state: receipt}
}

func (receipt *BookReceipt) State() data.BookReceipt {
	return receipt.state.Clone()
}
