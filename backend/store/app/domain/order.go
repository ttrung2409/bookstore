package domain

import (
	module "store"
	"store/app/data"
	"store/utils"
)

type Order struct {
	data.Order
}

var orderRepository = module.Container.Get(utils.Nameof((*data.OrderRepository)(nil))).(data.OrderRepository)

func (Order) FindByStatus(status string) ([]Order, error) {
	dataOrders, err := orderRepository.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	var orders []Order
	for _, dataOrder := range dataOrders {
		orders = append(orders, Order{dataOrder})
	}

	return orders, nil
}
