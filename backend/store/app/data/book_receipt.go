package data

import (
	"time"

	"github.com/thoas/go-funk"
)

type BookReceipt struct {
	Id                    EntityId `gorm:"primaryKey"`
	Number                uint     `gorm:"autoIncrement"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Items                 []BookReceiptItem
	OnhandStockAdjustment StockAdjustment
}

func (r *BookReceipt) GetId() EntityId {
	return r.Id
}

func (r *BookReceipt) SetId(id EntityId) {
	r.Id = id
}

func (r *BookReceipt) Clone() *BookReceipt {
	return &BookReceipt{
		Id:     r.Id,
		Number: r.Number,
		Items: funk.Map(r.Items, func(item BookReceiptItem) BookReceiptItem {
			return item.Clone()
		}).([]BookReceiptItem),
		OnhandStockAdjustment: r.OnhandStockAdjustment.Clone(),
	}
}

type BookReceiptItem struct {
	Id            EntityId `gorm:"primaryKey"`
	BookReceiptId EntityId
	BookId        EntityId
	Book          Book `gorm:"foreignKey:Id"`
	Qty           int
}

func (item BookReceiptItem) GetId() EntityId {
	return item.Id
}

func (item BookReceiptItem) SetId(id EntityId) {
	item.Id = id
}

func (item BookReceiptItem) Clone() BookReceiptItem {
	return BookReceiptItem{
		Id:            item.Id,
		BookReceiptId: item.BookReceiptId,
		BookId:        item.BookId,
		Book:          item.Book,
		Qty:           item.Qty,
	}
}

type BookReceiptRepository interface {
	repositoryBase
	Get(id EntityId, tx Transaction) (*BookReceipt, error)
	Create(receipt *BookReceipt, tx Transaction) (EntityId, error)
}

type BookReceiptItemRepository interface {
	repositoryBase
}
