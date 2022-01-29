package data

import "time"

type OrderStatus string

const (
	OrderStatusQueued      OrderStatus = "Queued"
	OrderStatusAccepted    OrderStatus = "Accepted"
	OrderStatusReceiving   OrderStatus = "Receiving"
	OrderStatusStockFilled OrderStatus = "StockFilled"
	OrderStatusRejected    OrderStatus = "Rejected"
)

type Order struct {
	Id         EntityId `gorm:"primaryKey"`
	Number     string   `gorm:"autoIncrement"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CustomerId EntityId
	Status     OrderStatus
	Items      []*OrderItem `gorm:"foreignKey:OrderId"`
	Stock      Stock        `gorm:"-"`
	Customer   *Customer    `gorm:"foreignKey:Id"`
}

func (o *Order) GetId() EntityId {
	return o.Id
}

func (o *Order) SetId(id EntityId) {
	o.Id = id
}

func (o *Order) Clone() *Order {
	items := []*OrderItem{}
	for _, item := range o.Items {
		items = append(items, item.Clone())
	}

	return &Order{
		Id:         o.Id,
		Number:     o.Number,
		CustomerId: o.CustomerId,
		Status:     o.Status,
		Items:      items,
		Stock:      o.Stock.Clone(),
	}
}

type OrderItem struct {
	Id      EntityId `gorm:"primaryKey"`
	OrderId EntityId
	BookId  EntityId
	Book    *Book `gorm:"foreignKey:Id"`
	Qty     int
}

func (item *OrderItem) Clone() *OrderItem {
	return &OrderItem{
		Id:      item.Id,
		OrderId: item.OrderId,
		BookId:  item.BookId,
		Book:    item.Book,
		Qty:     item.Qty,
	}
}
