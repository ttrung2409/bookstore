package query

import (
	"time"
)

type Order struct {
	Id        string
	Number    string
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
