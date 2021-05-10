package data

import "time"

type BookReceipt struct {
	Id        EntityId `gorm:"primaryKey"`
	Number    uint     `gorm:"autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookReceiptItem struct {
	BookReceiptId EntityId `gorm:"primaryKey"`
	BookId        EntityId `gorm:"primaryKey"`
	Qty           int
}

type BookReceiptRepository interface {
	repositoryBase
}

type BookReceiptItemRepository interface {
	repositoryBase
}
