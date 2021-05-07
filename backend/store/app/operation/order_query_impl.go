package operation

import (
	module "store"
	"store/app/data"
	"store/app/domain"
	"store/utils"
)

var orderRepository = module.Container.Get(utils.Nameof((*data.OrderRepository)(nil))).(data.OrderRepository)

type orderQuery struct{}

func (Order) FindByStatus(statuses []string) ([]Order, error) {
	dataOrders, err := orderRepository.FindByStatus(statuses)
	if err != nil {
		return nil, err
	}

	var orders []Order
	for _, dataOrder := range dataOrders {
		orders = append(orders, Order{}.fromDataObject(dataOrder))
	}

	return orders, nil
}

func (s *orderQuery) Get(id string) (*Order, error) {
	order, err := domain.Order{}.Get(data.FromStringToEntityId(id))
	if err != nil {
		return nil, err
	}

	viewOrder := Order{}.fromDataObject(order.Order)

	return &viewOrder, nil
}
