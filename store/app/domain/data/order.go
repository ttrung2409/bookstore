package data

import "time"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusAccepted  OrderStatus = "Accepted"
	OrderStatusRejected  OrderStatus = "Rejected"
	OrderStatusCancelled OrderStatus = "Cancelled"
	OrderStatusDelivered OrderStatus = "Delivered"
)

type Order struct {
	Id         string `gorm:"primaryKey"`
	Number     string `gorm:"autoIncrement"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CustomerId string
	Status     OrderStatus
	Items      []OrderItem `gorm:"foreignKey:OrderId"`
	Stock      Stock       `gorm:"-"`
	Customer   *Customer   `gorm:"foreignKey:Id"`
}

func (o *Order) GetId() string {
	return o.Id
}

func (o *Order) SetId(id string) {
	o.Id = id
}

func (o *Order) Clone() Order {
	items := []OrderItem{}
	for _, item := range o.Items {
		items = append(items, item.Clone())
	}

	return Order{
		Id:         o.Id,
		Number:     o.Number,
		CustomerId: o.CustomerId,
		Status:     o.Status,
		Items:      items,
		Stock:      o.Stock.Clone(),
	}
}

type OrderItem struct {
	OrderId string
	BookId  string
	Book    *Book `gorm:"foreignKey:Id"`
	Qty     int
}

func (item OrderItem) Clone() OrderItem {
	return OrderItem{
		OrderId: item.OrderId,
		BookId:  item.BookId,
		Book:    item.Book,
		Qty:     item.Qty,
	}
}
