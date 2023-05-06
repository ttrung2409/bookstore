package domain

import (
	"fmt"
	"store/app/domain/events"
	"time"
)

type OrderStatus string

const (
	OrderStatusAccepted  OrderStatus = "Accepted"
	OrderStatusRejected  OrderStatus = "Rejected"
	OrderStatusCancelled OrderStatus = "Cancelled"
	OrderStatusDelivered OrderStatus = "Delivered"
)

type OrderData struct {
	Id                      string `gorm:"primaryKey"`
	Status                  OrderStatus
	Items                   []OrderItemData `gorm:"foreignKey:OrderId"`
	CustomerId              string
	CustomerName            string
	CustomerPhone           string
	CustomerDeliveryAddress string
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func (o OrderData) Clone() OrderData {
	items := []OrderItemData{}
	for _, item := range o.Items {
		items = append(items, item.Clone())
	}

	return OrderData{
		Id:         o.Id,
		CustomerId: o.CustomerId,
		Status:     o.Status,
		Items:      items,
	}
}

type OrderItemData struct {
	OrderId string
	BookId  string
	Book    BookData `gorm:"foreignKey:Id"`
	Qty     int
}

func (item OrderItemData) Clone() OrderItemData {
	return OrderItemData{
		OrderId: item.OrderId,
		BookId:  item.BookId,
		Book:    item.Book,
		Qty:     item.Qty,
	}
}

type Order struct {
	eventSource
	state OrderData
	stock *Stock
}

func (Order) New(order OrderData, stock StockData) *Order {
	cloned := order.Clone()
	if cloned.Id == "" {
		cloned.Id = NewId()
	}

	return &Order{
		eventSource: eventSource{pendingEvents: []Event{}},
		state:       order.Clone(),
		stock:       Stock{}.New(stock),
	}
}

func (order *Order) State() struct {
	OrderData
	Stock StockData
} {
	return struct {
		OrderData
		Stock StockData
	}{
		OrderData: order.state.Clone(),
		Stock:     order.stock.State(),
	}
}

func (order *Order) Accept() error {
	if !order.stock.enoughForOrder(*order) {
		order.pendingEvents = append(
			order.pendingEvents,
			&events.OrderRejected{OrderId: order.state.Id},
		)

		return ErrNotEnoughStock
	}

	order.state.Status = OrderStatusAccepted
	order.stock = order.stock.reserveForOrder(*order)

	order.pendingEvents = append(
		order.pendingEvents,
		&events.OrderAccepted{OrderId: order.state.Id},
	)

	return nil
}

func (order *Order) Deliver() error {
	if order.state.Status != OrderStatusAccepted {
		return fmt.Errorf(
			"order status is '%s', no delivery allowed",
			order.state.Status,
		)
	}

	order.stock = order.stock.decreaseByOrder(*order)

	return nil
}

func (order *Order) Cancel() error {
	if order.state.Status != OrderStatusAccepted {
		return fmt.Errorf(
			"order status is '%s', no cancellation allowed",
			order.state.Status,
		)
	}

	order.stock = order.stock.releaseReservation(*order)

	return nil
}
