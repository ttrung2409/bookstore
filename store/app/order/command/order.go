package command

import (
	"store/app/domain/data"
	"time"

	"github.com/thoas/go-funk"
)

type Order struct {
	Id        string
	CreatedAt time.Time
	Customer  Customer
	Status    string
	Items     []OrderItem
}

func (o *Order) toDataObject() data.Order {
	return data.Order{
		Id:                      o.Id,
		Status:                  data.OrderStatus(o.Status),
		CustomerId:              o.Customer.Id,
		CustomerName:            o.Customer.Name,
		CustomerPhone:           o.Customer.Phone,
		CustomerDeliveryAddress: o.Customer.DeliveryAddress,
		Items: funk.Map(o.Items, func(item OrderItem) data.OrderItem {
			return data.OrderItem{OrderId: o.Id, BookId: item.BookId, Qty: item.Qty}
		}).([]data.OrderItem),
	}
}

type OrderItem struct {
	BookId string
	Qty    int
}

type Customer struct {
	Id              string
	Name            string
	Phone           string
	DeliveryAddress string
}
