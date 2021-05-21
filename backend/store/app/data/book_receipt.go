package data

import "time"

type BookReceipt struct {
	Id        EntityId `gorm:"primaryKey"`
	Number    uint     `gorm:"autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []BookReceiptItem
	Stock     Stock `gorm:"-"`
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
	Book          Book `gorm:"foreignKey:Id"`
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
	Get(id EntityId, tx Transaction) (*BookReceipt, error)
	Create(receipt *BookReceipt, tx Transaction) (EntityId, error)
	Update(receipt *BookReceipt, tx Transaction) error
}

type BookReceiptItemRepository interface {
	repositoryBase
}
