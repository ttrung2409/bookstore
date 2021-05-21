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
	Stock      Stock       `gorm:"-"`
}

func (o *Order) GetId() EntityId {
	return o.Id
}

func (o *Order) SetId(id EntityId) {
	o.Id = id
}

type OrderItem struct {
	Id      EntityId `gorm:"primaryKey"`
	OrderId EntityId
	BookId  EntityId
	Book    Book `gorm:"foreignKey:Id"`
	Qty     int
}

type OrderRepository interface {
	repositoryBase
	Get(id EntityId, tx Transaction) (*Order, error)
	GetReceivingOrders(tx Transaction) ([]*Order, error)
	Update(order *Order, tx Transaction) error
}
