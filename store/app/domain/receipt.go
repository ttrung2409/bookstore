package domain

import (
	data "store/app/domain/data"

	"github.com/thoas/go-funk"
)

type Receipt struct {
	state data.Receipt
}

type ReceivingBook struct {
	data.Book
	ReceivingQty int
}

func (Receipt) New(state data.Receipt) *Receipt {
	cloned := state.Clone()
	if cloned.Id == "" {
		cloned.Id = data.NewId()
	}

	receipt := &Receipt{state: cloned}
	return receipt
}

func (Receipt) NewFromReceivingBooks(books []ReceivingBook) *Receipt {
	receipt := data.Receipt{
		Id: data.NewId(),
	}

	items := []data.ReceiptItem{}

	for _, book := range books {
		items = append(items, data.ReceiptItem{
			ReceiptId: receipt.Id,
			BookId:    book.Id,
			Qty:       book.ReceivingQty,
		})
	}

	receipt.Items = items
	receipt.OnhandStockAdjustment = funk.Map(receipt.Items, func(item data.ReceiptItem) data.StockAdjustmentItem {
		return data.StockAdjustmentItem{BookId: item.BookId, Qty: item.Qty}
	}).(data.StockAdjustment)

	return &Receipt{state: receipt}
}

func (receipt *Receipt) State() data.Receipt {
	return receipt.state.Clone()
}
