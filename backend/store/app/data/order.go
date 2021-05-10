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
	Id         EntityId `gorm:"primaryKey"`
	Number     string   `gorm:"autoIncrement"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CustomerId EntityId
	Status     OrderStatus
	Items      []OrderItem `gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	OrderId EntityId `gorm:"primaryKey"`
	BookId  EntityId `gorm:"primaryKey"`
	Book    Book     `gorm:"foreignKey:Id"`
	Qty     int
}

type OrderRepository interface {
	repositoryBase
}
