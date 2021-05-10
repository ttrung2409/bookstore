package domain

import (
	"errors"
	"fmt"
	module "store"
	"store/app/data"
	"store/utils"
)

type Order struct {
	data.Order
}

var orderRepository = module.Container().Get(utils.Nameof((*data.OrderRepository)(nil))).(data.OrderRepository)

func (Order) Get(id data.EntityId) (*Order, error) {
	result, err := orderRepository.
		Query(&data.Order{}).
		IncludeMany("Items").
		Where("id = ?", id).
		First()

	if err != nil {
		return nil, err
	}

	dataOrder := result.(data.Order)

	return &Order{dataOrder}, nil
}

func (order *Order) Accept() error {
	if order.Status != data.OrderStatusQueued && order.Status != data.OrderStatusStockFilled {
		return errors.New(fmt.Sprintf("Order status '%s' is invalid for this operation", order.Status))
	}

	tx := transactionFactory.New()

	orderRepository.Update(
		order.Id,
		&data.Order{Status: data.OrderStatusAccepted},
		tx,
	)

	err := tx.Commit()

	if err != nil {
		tx.Rollback()

		return err
	}

	return nil
}
