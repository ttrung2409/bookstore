package domain

import (
	"errors"
	"fmt"
	"store/app/data"
)

type Order struct {
	data.Order
}

func (Order) Get(id data.EntityId, tx data.Transaction) (*Order, error) {
	result, err := OrderRepository.
		Query(&data.Order{}, tx).
		IncludeMany("Items").
		Where("id = ?", id).
		First()

	if err != nil {
		return nil, err
	}

	dataOrder := result.(data.Order)

	return &Order{dataOrder}, nil
}

func (Order) GetReceivingOrders(tx data.Transaction) ([]*Order, error) {
	records, err := OrderRepository.
		Query(&data.Order{}, tx).
		Where("status = ?", data.OrderStatusReceiving).
		IncludeMany("Items").
		OrderBy("created_at").
		Find()

	if err != nil {
		return nil, err
	}

	orders := []*Order{}
	for _, record := range records {
		orders = append(orders, &Order{record.(data.Order)})
	}

	return orders, nil
}

func (order *Order) Accept(tx data.Transaction) error {
	if order.Status != data.OrderStatusQueued && order.Status != data.OrderStatusStockFilled {
		return errors.New(fmt.Sprintf("Order status '%s' is invalid for accepting", order.Status))
	}

	err := OrderRepository.Update(
		order.Id,
		&data.Order{Status: data.OrderStatusAccepted},
		tx,
	)

	if err != nil {
		return err
	}

	return nil
}

func (order *Order) TryUpdateToStockFilled(
	stock data.Stock,
	tx data.Transaction,
) (data.Stock, error) {
	if order.Status != data.OrderStatusReceiving {
		return stock, errors.New(
			fmt.Sprintf("Order status '%s' is invalid for StockFilled", order.Status),
		)
	}

	if !stock.Enough(order.Items) {
		return stock, errors.New("Not enough stock")
	}

	err := OrderRepository.Update(
		order.Id,
		&data.Order{Status: data.OrderStatusStockFilled},
		tx,
	)

	if err != nil {
		return stock, err
	}

	return stock.Issue(order.Items), nil
}
