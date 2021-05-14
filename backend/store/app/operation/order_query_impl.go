package operation

import (
	module "store"
	"store/app/data"
	"store/utils"
)

var orderRepository = module.Container().Get(utils.Nameof((*data.OrderRepository)(nil))).(data.OrderRepository)

type orderQuery struct{}

func (*orderQuery) FindByStatus(statuses []string) ([]Order, error) {
	records, err := orderRepository.Query(&data.Order{}, nil).Where("status IN ?", statuses).Find()
	if err != nil {
		return nil, err
	}

	var orders []Order
	for _, record := range records {
		dataOrder := record.(data.Order)
		orders = append(orders, Order{}.fromDataObject(dataOrder))
	}

	return orders, nil
}

func (*orderQuery) GetWithItems(id string) (*Order, error) {
	orderId := data.FromStringToEntityId(id)
	result, err := orderRepository.
		Query(&data.Order{}, nil).
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id = ?", orderId).
		First()

	if err != nil {
		return nil, err
	}

	order := Order{}.fromDataObject(result.(data.Order))

	return &order, nil
}
