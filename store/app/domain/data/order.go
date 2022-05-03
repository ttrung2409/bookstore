package data

import (
	"fmt"
	"time"
)

type OrderStatus string

const (
	OrderStatusAccepted  OrderStatus = "Accepted"
	OrderStatusRejected  OrderStatus = "Rejected"
	OrderStatusCancelled OrderStatus = "Cancelled"
	OrderStatusDelivered OrderStatus = "Delivered"
)

type Order struct {
	Id                      string `gorm:"primaryKey"`
	Status                  OrderStatus
	Items                   []OrderItem `gorm:"foreignKey:OrderId"`
	Stock                   Stock       `gorm:"-"`
	CustomerId              string
	CustomerName            string
	CustomerPhone           string
	CustomerDeliveryAddress string
	CreatedAt               time.Time
	UpdatedAt               time.Time
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

func (item OrderItem) GetId() string {
	return fmt.Sprintf("%s-%s", item.OrderId, item.BookId)
}

func (item OrderItem) SetId(id string) {
}

func (item OrderItem) Clone() OrderItem {
	return OrderItem{
		OrderId: item.OrderId,
		BookId:  item.BookId,
		Book:    item.Book,
		Qty:     item.Qty,
	}
}
