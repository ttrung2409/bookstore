package query

import "time"

type Order struct {
	Id        string
	Number    string
	Status    string
	CreatedAt time.Time
	Customer  Customer
	Items     []OrderItem
}

type OrderItem struct {
	Book Book
	Qty  int
}
