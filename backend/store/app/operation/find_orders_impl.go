package operation

import (
	"store/app/domain"
)

type findOrders struct{}

func (s *findOrders) Find(status string) ([]Order, error) {
	orders, err := domain.Order{}.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	var viewOrders []Order
	for _, order := range orders {
		viewOrders = append(viewOrders, Order{}.fromDomainObject(order))
	}

	return viewOrders, nil
}
