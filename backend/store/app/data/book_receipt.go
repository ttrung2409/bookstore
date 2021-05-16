package data

import "time"

type BookReceipt struct {
	Id        EntityId `gorm:"primaryKey"`
	Number    uint     `gorm:"autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []BookReceiptItem
}

func (r *BookReceipt) GetId() EntityId {
	return r.Id
}

func (r *BookReceipt) SetId(id EntityId) {
	r.Id = id
}

type BookReceiptItem struct {
	Id            EntityId `gorm:"primaryKey"`
	BookReceiptId EntityId
	BookId        EntityId
	Qty           int
}

func (i *BookReceiptItem) GetId() EntityId {
	return i.Id
}

func (i *BookReceiptItem) SetId(id EntityId) {
	i.Id = id
}

type BookReceiptRepository interface {
	repositoryBase
}

type BookReceiptItemRepository interface {
	repositoryBase
}
