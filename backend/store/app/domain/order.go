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
	Items []data.OrderItem
}

var orderRepository = module.Container.Get(utils.Nameof((*data.OrderRepository)(nil))).(data.OrderRepository)

var orderItemRepository = module.Container.Get(utils.Nameof((*data.OrderItemRepository)(nil))).(data.OrderItemRepository)

func (Order) Get(id data.EntityId) (*Order, error) {
	dataOrder, err := orderRepository.Get(id)
	if err != nil {
		return nil, err
	}

	dataOrder, ok := dataOrder.(data.Order)
	if !ok {
		return nil, nil
	}

	items, err := orderItemRepository.GetByOrderId(id)

	return &Order{dataOrder, Items: items}, nil
}

func (o *Order) Accept() error {
	if o.Status != data.OrderStatusQueued && o.Status != data.OrderStatusStockFilled {
		return errors.New(fmt.Sprintf("Order status '%s' is invalid for this operation", o.Status))
	}

	transaction := transactionFactory.New()

	orderRepository.Update(
		o.Id,
		struct{ status data.OrderStatus }{status: data.OrderStatusAccepted},
		transaction,
	)

	err := (*transaction).Commit()

	if err != nil {
		(*transaction).Rollback()

		return err
	}

	return nil
}
