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

func (r *BookReceipt) Clone() *BookReceipt {
	items := []BookReceiptItem{}
	for _, item := range r.Items {
		items = append(items, item.Clone())
	}

	return &BookReceipt{
		Id:     r.Id,
		Number: r.Number,
		Items:  items,
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
	Update(receipt *BookReceipt, tx Transaction) error
}

type BookReceiptItemRepository interface {
	repositoryBase
}
