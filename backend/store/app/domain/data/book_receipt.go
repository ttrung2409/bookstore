package data

import (
	"time"

	"github.com/thoas/go-funk"
)

type BookReceipt struct {
	Id                    string `gorm:"primaryKey"`
	Number                uint   `gorm:"autoIncrement"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Items                 []*BookReceiptItem
	OnhandStockAdjustment StockAdjustment
}

func (r *BookReceipt) GetId() string {
	return r.Id
}

func (r *BookReceipt) SetId(id string) {
	r.Id = id
}

func (r *BookReceipt) Clone() *BookReceipt {
	return &BookReceipt{
		Id:     r.Id,
		Number: r.Number,
		Items: funk.Map(r.Items, func(item *BookReceiptItem) *BookReceiptItem {
			return item.Clone()
		}).([]*BookReceiptItem),
		OnhandStockAdjustment: r.OnhandStockAdjustment.Clone(),
	}
}

type BookReceiptItem struct {
	Id            string `gorm:"primaryKey"`
	BookReceiptId string
	BookId        string
	Book          *Book `gorm:"foreignKey:Id"`
	Qty           int
}

func (item BookReceiptItem) GetId() string {
	return item.Id
}

func (item BookReceiptItem) SetId(id string) {
	item.Id = id
}

func (item *BookReceiptItem) Clone() *BookReceiptItem {
	return &BookReceiptItem{
		Id:            item.Id,
		BookReceiptId: item.BookReceiptId,
		BookId:        item.BookId,
		Book:          item.Book,
		Qty:           item.Qty,
	}
}
