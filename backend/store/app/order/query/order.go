package query

import (
	"time"
)

type Order struct {
	Id         string
	Number     string
	CreatedAt  time.Time
	CustomerId string
	Status     string
	Items      []OrderItem
}

type OrderItem struct {
	Book Book
	Qty  int
}
