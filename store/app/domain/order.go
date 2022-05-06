package domain

import (
	"fmt"
	"store/app/domain/data"
	"store/app/domain/events"
)

type Order struct {
	EventSource
	state data.Order
	stock *Stock
}

func (Order) New(state data.Order) *Order {
	return &Order{
		EventSource: EventSource{PendingEvents: []Event{}},
		state:       state,
		stock:       Stock{}.New(state.Stock),
	}
}

func (order *Order) State() data.Order {
	return order.state.Clone()
}

func (order *Order) Accept() error {
	if !order.stock.enoughForOrder(*order) {
		order.PendingEvents = append(
			order.PendingEvents,
			&events.OrderRejected{OrderId: order.state.Id},
		)

		return ErrNotEnoughStock
	}

	order.state.Status = data.OrderStatusAccepted
	order.stock = order.stock.reserveForOrder(*order)

	order.PendingEvents = append(
		order.PendingEvents,
		&events.OrderAccepted{OrderId: order.state.Id},
	)

	return nil
}

func (order *Order) Deliver() error {
	if order.state.Status != data.OrderStatusAccepted {
		return fmt.Errorf(
			"order status '%s' is invalid. Order status must be 'Accepted' so as for it to be delivered",
			order.state.Status,
		)
	}

	order.stock = order.stock.decreaseByOrder(*order)

	return nil
}

func (order *Order) Cancel() error {
	if order.state.Status != data.OrderStatusAccepted {
		return fmt.Errorf(
			"order status '%s' is invalid. Order status must be 'Accepted' so as for it to be cancelled",
			order.state.Status,
		)
	}

	order.stock = order.stock.releaseReservation(*order)

	return nil
}
