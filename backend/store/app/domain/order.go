package domain

import (
	"errors"
	"fmt"
	"store/app/data"
)

type Order struct {
	data.Order
}

func (Order) New(dataOrder *data.Order) *Order {
	return &Order{Order: *dataOrder}
}

func (order *Order) Accept(tx data.Transaction) error {
	if order.Status != data.OrderStatusQueued && order.Status != data.OrderStatusStockFilled {
		return errors.New(fmt.Sprintf("Order status '%s' is invalid for accepting", order.Status))
	}

	order.Status = data.OrderStatusAccepted

	stock := Stock{Stock: order.Stock}
	if !stock.EnoughForOrder(order) {
		return errors.New("Not enough stock")
	}

	order.Stock = stock.DecreaseByOrder(order).Stock

	return nil
}

func (order *Order) PlaceAsBackOrder(tx data.Transaction) error {
	if order.Status != data.OrderStatusQueued {
		return errors.New(
			fmt.Sprintf("Order status '%s' is invalid to be placed as backorder", order.Status),
		)
	}

	order.Status = data.OrderStatusReceiving

	stock := Stock{Stock: order.Stock}
	order.Stock = stock.ReserveForOrder(order).Stock

	return nil
}

func (order *Order) TryUpdateToStockFilled(
	stock Stock,
	tx data.Transaction,
) (Stock, error) {
	if order.Status != data.OrderStatusReceiving {
		return stock, errors.New(
			fmt.Sprintf("Order status '%s' is invalid for StockFilled", order.Status),
		)
	}

	if !stock.EnoughForOrder(order) {
		return stock, errors.New("Not enough stock")
	}

	order.Status = data.OrderStatusStockFilled
	order.Stock = stock.Clone().Stock

	return stock.DecreaseByOrder(order), nil
}

func (order *Order) Reject(tx data.Transaction) error {
	if order.Status != data.OrderStatusQueued && order.Status != data.OrderStatusStockFilled {
		return errors.New(fmt.Sprintf("Order status '%s' is invalid for rejecting", order.Status))
	}

	order.Status = data.OrderStatusRejected

	return nil
}
