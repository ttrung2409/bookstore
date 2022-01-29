package domain

import (
	"errors"
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
	if order.state.Status != data.OrderStatusQueued &&
		order.state.Status != data.OrderStatusStockFilled {
		return fmt.Errorf("Order status '%s' is invalid for accepting", order.state.Status)
	}

	stock := Stock{}.New(order.state.Stock)
	if !stock.enoughForOrder(order) {
		return errors.New("Not enough stock")
	}

	order.state.Status = data.OrderStatusAccepted
	order.state.Stock = stock.decreaseByOrder(order).state

	return nil
}

func (order *Order) PlaceAsBackOrder() error {
	if order.state.Status != data.OrderStatusQueued &&
		order.state.Status != data.OrderStatusStockFilled {
		return fmt.Errorf("Order status '%s' is invalid to be placed as backorder", order.state.Status)
	}

	order.state.Status = data.OrderStatusReceiving

	stock := Stock{}.New(order.state.Stock)
	order.state.Stock = stock.reserveForOrder(order).state

	return nil
}

func (order *Order) UpdateToStockFilled() (bool, error) {
	if order.state.Status != data.OrderStatusReceiving {
		return false, fmt.Errorf("Order status '%s' is invalid for StockFilled", order.state.Status)
	}

	stock := Stock{}.New(order.state.Stock)
	if !stock.enoughForOrder(order) {
		return false, errors.New("Not enough stock")
	}

	order.state.Stock = stock.releaseReservation(order).state
	order.state.Status = data.OrderStatusStockFilled

	return true, nil
}

func (order *Order) Reject() error {
	if order.state.Status == data.OrderStatusAccepted ||
		order.state.Status == data.OrderStatusRejected {
		return fmt.Errorf("Order status '%s' is invalid for rejecting", order.state.Status)
	}

	stock := Stock{}.New(order.state.Stock)
	if order.state.Status == data.OrderStatusReceiving {
		order.state.Stock = stock.releaseReservation(order).state
	}

	order.state.Status = data.OrderStatusRejected

	return nil
}
