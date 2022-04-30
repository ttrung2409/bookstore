package query

import (
	"store/app/domain/data"
	"time"
)

type Order struct {
	Id        string
	Number    string
	CreatedAt time.Time
	Customer  *Customer
	Status    string
	Items     []*OrderItem
}

type OrderItem struct {
	Book *Book
	Qty  int
}

type Customer struct {
	Name            string
	Phone           string
	DeliveryAddress string
}

func (Order) fromDataObject(order *data.Order) *Order {
	items := []*OrderItem{}
	for _, dataItem := range order.Items {
		items = append(items, OrderItem{}.fromDataObject(dataItem))
	}

	return &Order{
		Id:        order.Id,
		Number:    order.Number,
		CreatedAt: order.CreatedAt,
		Customer:  Customer{}.fromDataObject(order.Customer),
		Status:    string(order.Status),
		Items:     items,
	}
}

func (OrderItem) fromDataObject(item *data.OrderItem) *OrderItem {
	return &OrderItem{
		Book: Book{}.fromDataObject(item.Book),
		Qty:  item.Qty,
	}
}

func (Customer) fromDataObject(customer *data.Customer) *Customer {
	return &Customer{
		Name:            customer.Name,
		Phone:           customer.Phone,
		DeliveryAddress: customer.DeliveryAddress,
	}
}
