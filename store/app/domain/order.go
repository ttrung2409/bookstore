package domain

import (
	"fmt"
	"store/app/domain/data"
)

type Order struct {
	state *data.Order
}

func (Order) New(state *data.Order) *Order {
	return &Order{state: state}
}

func (order *Order) State() *data.Order {
	return order.state.Clone()
}

func (order *Order) Accept() error {
	if order.state.Status != data.OrderStatusPending {
		return fmt.Errorf("order status '%s' is invalid. Order status must be 'Pending' so as for it to be accepted", order.state.Status)
	}

	stock := Stock{}.New(order.state.Stock)
	if !stock.enoughForOrder(order) {
		return ErrNotEnoughStock
	}

	order.state.Status = data.OrderStatusAccepted
	order.state.Stock = stock.reserveForOrder(order).state

	return nil
}

func (order *Order) Deliver() error {
	if order.state.Status != data.OrderStatusAccepted {
		return fmt.Errorf("order status '%s' is invalid. Order status must be 'Accepted' so as for it to be delivered", order.state.Status)
	}

	stock := Stock{}.New(order.state.Stock)
	order.state.Stock = stock.decreaseByOrder(order).state

	return nil
}

func (order *Order) Cancel() error {
	if order.state.Status != data.OrderStatusAccepted {
		return fmt.Errorf("order status '%s' is invalid. Order status must be 'Accepted' so as for it to be cancelled", order.state.Status)
	}

	stock := Stock{}.New(order.state.Stock)
	order.state.Stock = stock.releaseReservation(order).state

	return nil
}
