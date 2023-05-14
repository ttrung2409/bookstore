package domain

import (
	"fmt"
	"store/app/domain/events"
	"time"

	"github.com/thoas/go-funk"
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
	Items                   []OrderItem `gorm:"foreignKey:OrderId"`
	CustomerId              string
	CustomerName            string
	CustomerPhone           string
	CustomerDeliveryAddress string
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func (o OrderData) Clone() OrderData {
	items := []OrderItem{}
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

type OrderItem struct {
	OrderId string
	BookId  string
	Book    BookData `gorm:"foreignKey:Id"`
	Qty     int
}

func (item OrderItem) Clone() OrderItem {
	return OrderItem{
		OrderId: item.OrderId,
		BookId:  item.BookId,
		Book:    item.Book,
		Qty:     item.Qty,
	}
}

type Order struct {
	eventSource
	order           OrderData
	stock           Stock
	stockAdjustment StockAdjustment
}

func (Order) New(order OrderData, stock Stock) *Order {
	cloned := order.Clone()
	if cloned.Id == "" {
		cloned.Id = NewId()
	}

	return &Order{
		eventSource: eventSource{pendingEvents: []Event{}},
		order:       order.Clone(),
	}
}

func (order *Order) State() struct {
	OrderData
	StockAdjustment StockAdjustment
} {
	return struct {
		OrderData
		StockAdjustment StockAdjustment
	}{
		OrderData:       order.order.Clone(),
		StockAdjustment: order.stockAdjustment.Clone(),
	}
}

func (order *Order) Accept() error {
	stockEnoughForOrder := true

	for _, item := range order.order.Items {
		if stockItem, ok := order.stock[item.BookId]; ok {
			if item.Qty > stockItem.OnhandQty-stockItem.ReservedQty {
				stockEnoughForOrder = false
			}
		}
	}

	if !stockEnoughForOrder {
		order.pendingEvents = append(
			order.pendingEvents,
			&events.OrderRejected{OrderId: order.order.Id},
		)

		return ErrNotEnoughStock
	}

	order.order.Status = OrderStatusAccepted
	order.stockAdjustment = funk.Map(order.order.Items, func(item OrderItem) StockAdjustmentItem {
		return StockAdjustmentItem{BookId: item.BookId, Qty: item.Qty, StockType: StockTypeReserved}
	}).(StockAdjustment)

	order.pendingEvents = append(
		order.pendingEvents,
		&events.OrderAccepted{OrderId: order.order.Id},
	)

	return nil
}

func (order *Order) Deliver() error {
	if order.order.Status != OrderStatusAccepted {
		return fmt.Errorf(
			"order status is '%s', no delivery allowed",
			order.order.Status,
		)
	}

	order.order.Status = OrderStatusDelivered
	order.stockAdjustment = funk.Map(order.order.Items, func(item OrderItem) StockAdjustmentItem {
		return StockAdjustmentItem{BookId: item.BookId, Qty: -item.Qty, StockType: StockTypeOnhand}
	}).(StockAdjustment)

	order.pendingEvents = append(
		order.pendingEvents,
		&events.OrderDelivered{OrderId: order.order.Id},
	)

	return nil
}

func (order *Order) Cancel() error {
	if order.order.Status != OrderStatusAccepted {
		return fmt.Errorf(
			"order status is '%s', no cancellation allowed",
			order.order.Status,
		)
	}

	order.order.Status = OrderStatusCancelled
	order.stockAdjustment = funk.Map(order.order.Items, func(item OrderItem) StockAdjustmentItem {
		return StockAdjustmentItem{BookId: item.BookId, Qty: -item.Qty, StockType: StockTypeReserved}
	}).(StockAdjustment)

	return nil
}
