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
	Items     []ReceiptItemData `gorm:"foreignKey:ReceiptId"`
}

func (r ReceiptData) Clone() ReceiptData {
	return ReceiptData{
		Id: r.Id,
		Items: funk.Map(r.Items, func(item ReceiptItemData) ReceiptItemData {
			return item.Clone()
		}).([]ReceiptItemData),
	}
}

type ReceiptItemData struct {
	ReceiptId string
	BookId    string
	Book      Book `gorm:"foreignKey:Id"`
	Qty       int
}

func (item ReceiptItemData) Clone() ReceiptItemData {
	return ReceiptItemData{
		ReceiptId: item.ReceiptId,
		BookId:    item.BookId,
		Book:      item.Book,
		Qty:       item.Qty,
	}
}

type Receipt struct {
	state                 ReceiptData
	onhandStockAdjustment StockAdjustmentData
}

func (Receipt) New(receipt ReceiptData) *Receipt {
	cloned := receipt.Clone()
	if cloned.Id == "" {
		cloned.Id = NewId()
	}

	return &Receipt{
		state: cloned,
		onhandStockAdjustment: funk.Map(receipt.Items, func(item ReceiptItemData) StockAdjustmentItemData {
			return StockAdjustmentItemData{BookId: item.BookId, Qty: item.Qty}
		}).(StockAdjustmentData),
	}
}

func (Receipt) NewFromReceivingBooks(books []ReceivingBook) *Receipt {
	receipt := ReceiptData{
		Id: NewId(),
	}

	items := []ReceiptItemData{}

	for _, book := range books {
		items = append(items, ReceiptItemData{
			ReceiptId: receipt.Id,
			BookId:    book.Id,
			Qty:       book.ReceivingQty,
		})
	}

	receipt.Items = items

	return &Receipt{
		state: receipt,
		onhandStockAdjustment: funk.Map(receipt.Items, func(item ReceiptItemData) StockAdjustmentItemData {
			return StockAdjustmentItemData{BookId: item.BookId, Qty: item.Qty}
		}).(StockAdjustmentData),
	}
}

func (receipt *Receipt) State() struct {
	ReceiptData
	OnhandStockAdjustment StockAdjustmentData
} {
	return struct {
		ReceiptData
		OnhandStockAdjustment StockAdjustmentData
	}{
		ReceiptData:           receipt.state.Clone(),
		OnhandStockAdjustment: receipt.onhandStockAdjustment.Clone(),
	}
}
