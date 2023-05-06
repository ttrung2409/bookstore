package data

import (
	"time"

	"github.com/thoas/go-funk"
)

type Receipt struct {
	Id                    string `gorm:"primaryKey"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Items                 []ReceiptItem   `gorm:"foreignKey:ReceiptId"`
	OnhandStockAdjustment StockAdjustment `gorm:"-"`
}

func (r Receipt) Clone() Receipt {
	return Receipt{
		Id: r.Id,
		Items: funk.Map(r.Items, func(item ReceiptItem) ReceiptItem {
			return item.Clone()
		}).([]ReceiptItem),
		OnhandStockAdjustment: r.OnhandStockAdjustment.Clone(),
	}
}

type ReceiptItem struct {
	ReceiptId string
	BookId    string
	Book      Book `gorm:"foreignKey:Id"`
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
