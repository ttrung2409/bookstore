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
	Id         EntityId
	Number     string
	CreatedAt  time.Time
	CustomerId EntityId
	Status     OrderStatus
}

type OrderItem struct {
	OrderId EntityId
	BookId  BookId
	Qty     int
}

type OrderRepository interface {
	repositoryBase
	FindByStatus(statuses []string) ([]Order, error)
}

type OrderItemRepository interface {
	repositoryBase
	GetByOrderId(orderId EntityId) ([]OrderItem, error)
}
