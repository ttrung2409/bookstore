package data

import (
	"fmt"
	"time"

	"github.com/thoas/go-funk"
)

type BookReceipt struct {
	Id                    string `gorm:"primaryKey"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Items                 []BookReceiptItem `gorm:"foreignKey:BookReceiptId"`
	OnhandStockAdjustment StockAdjustment   `gorm:"-"`
}

func (r *BookReceipt) Clone() BookReceipt {
	return BookReceipt{
		Id: r.Id,
		Items: funk.Map(r.Items, func(item BookReceiptItem) BookReceiptItem {
			return item.Clone()
		}).([]BookReceiptItem),
		OnhandStockAdjustment: r.OnhandStockAdjustment.Clone(),
	}
}

type BookReceiptItem struct {
	BookReceiptId string
	BookId        string
	Book          *Book `gorm:"foreignKey:Id"`
	Qty           int
}

func (item *BookReceiptItem) GetId() string {
	return fmt.Sprintf("%s-%s", item.BookReceiptId, item.BookId)
}

func (item *BookReceiptItem) SetId(id string) {
}

func (item *BookReceiptItem) Clone() BookReceiptItem {
	return BookReceiptItem{
		BookReceiptId: item.BookReceiptId,
		BookId:        item.BookId,
		Book:          item.Book,
		Qty:           item.Qty,
	}
}
