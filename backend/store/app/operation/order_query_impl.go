package operation

import (
	module "store"
	"store/app/data"
	"store/utils"
)

var orderRepository = module.Container().Get(utils.Nameof((*data.OrderRepository)(nil))).(data.OrderRepository)

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

func (s *orderQuery) GetWithItems(id string) (*Order, error) {
	orderId := data.FromStringToEntityId(id)
	result, err := orderRepository.
		Query(&data.Order{}).
		Include("Items").
		ThenInclude("Book").
		Where("id = ?", orderId).
		First()

	if err != nil {
		return nil, err
	}

	order := Order{}.fromDataObject(result.(data.Order))

	return &order, nil
}
