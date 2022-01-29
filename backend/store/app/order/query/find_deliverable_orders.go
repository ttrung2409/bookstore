package query

import (
	"store/app/domain/data"
)

func (*query) FindDeliverableOrders() ([]*Order, error) {
	records, err := OrderRepository.
		Query(&data.Order{}, nil).
		Where("status IN ?",
			[]string{string(data.OrderStatusQueued), string(data.OrderStatusStockFilled)}).
		Find()

	if err != nil {
		return nil, err
	}

	var orders []*Order
	for _, record := range records {
		dataOrder := record.(*data.Order)
		viewOrder := Order{}.fromDataObject(dataOrder)
		orders = append(orders, &viewOrder)
	}

	return orders, nil
}
