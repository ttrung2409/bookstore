package data

import "time"

type OrderStatus int

const (
	Queued      OrderStatus = iota
	Accepted    OrderStatus = iota
	Receiving   OrderStatus = iota
	StockFilled OrderStatus = iota
	Rejected    OrderStatus = iota
)

type Order struct {
	Id         EntityId
	Number     string
	CreatedAt  time.Time
	CustomerId EntityId
	Status     OrderStatus
}
