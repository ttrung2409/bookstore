package domain

import (
	"time"

	"github.com/thoas/go-funk"
)

type ReceivingBook struct {
	BookData
	ReceivingQty int
}

type ReceiptData struct {
	Id        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []ReceiptItem `gorm:"foreignKey:ReceiptId"`
}

func (r ReceiptData) Clone() ReceiptData {
	return ReceiptData{
		Id: r.Id,
		Items: funk.Map(r.Items, func(item ReceiptItem) ReceiptItem {
			return item.Clone()
		}).([]ReceiptItem),
	}
}

type ReceiptItem struct {
	ReceiptId string
	BookId    string
	Book      BookData `gorm:"foreignKey:Id"`
	Qty       int
}

func (item ReceiptItem) Clone() ReceiptItem {
	return ReceiptItem{
		ReceiptId: item.ReceiptId,
		BookId:    item.BookId,
		Book:      item.Book,
		Qty:       item.Qty,
	}
}

type Receipt struct {
	state           ReceiptData
	stockAdjustment StockAdjustment
}

func (Receipt) NewFromReceivingBooks(books []ReceivingBook) *Receipt {
	receipt := ReceiptData{
		Id: NewId(),
	}

	items := []ReceiptItem{}

	for _, book := range books {
		items = append(items, ReceiptItem{
			ReceiptId: receipt.Id,
			BookId:    book.Id,
			Qty:       book.ReceivingQty,
		})
	}

	receipt.Items = items

	return &Receipt{
		state: receipt,
		stockAdjustment: funk.Map(receipt.Items, func(item ReceiptItem) StockAdjustmentItem {
			return StockAdjustmentItem{BookId: item.BookId, Qty: item.Qty}
		}).(StockAdjustment),
	}
}

func (receipt *Receipt) State() struct {
	ReceiptData
	StockAdjustment StockAdjustment
} {
	return struct {
		ReceiptData
		StockAdjustment StockAdjustment
	}{
		ReceiptData:     receipt.state.Clone(),
		StockAdjustment: receipt.stockAdjustment.Clone(),
	}
}
