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
	result, err := orderRepository.Query(&data.Order{}).Include("Items").Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	dataOrder := result.(data.Order)

	return &Order{dataOrder}, nil
}

func (o *Order) Accept() error {
	if o.Status != data.OrderStatusQueued && o.Status != data.OrderStatusStockFilled {
		return errors.New(fmt.Sprintf("Order status '%s' is invalid for this operation", o.Status))
	}

	tx := transactionFactory.New()

	orderRepository.Update(
		o.Id,
		struct{ status data.OrderStatus }{status: data.OrderStatusAccepted},
		tx,
	)

	err := tx.Commit()

	if err != nil {
		tx.Rollback()

		return err
	}

	return nil
}
