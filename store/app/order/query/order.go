package query

import (
	"store/app/domain"
	"time"
)

type Order struct {
	Id        string
	CreatedAt time.Time
	Customer  Customer
	Status    string
	Items     []OrderItem
}

type OrderItem struct {
	Book Book
	Qty  int
}

type Customer struct {
	Name            string
	Phone           string
	DeliveryAddress string
}

func (Order) fromDataObject(order domain.OrderData) Order {
	items := []OrderItem{}
	for _, dataItem := range order.Items {
		items = append(items, OrderItem{}.fromDataObject(dataItem))
	}

	return Order{
		Id:        order.Id,
		CreatedAt: order.CreatedAt,
		Customer: Customer{
			Name:            order.CustomerName,
			Phone:           order.CustomerPhone,
			DeliveryAddress: order.CustomerDeliveryAddress,
		},
		Status: string(order.Status),
		Items:  items,
	}
}

func (OrderItem) fromDataObject(item domain.OrderItemData) OrderItem {
	return OrderItem{
		Book: Book{}.fromDataObject(item.Book),
		Qty:  item.Qty,
	}
}
