package query

import (
	"store/app/data"
)

type orderQuery struct{}

func (*orderQuery) FindOrdersToDeliver() ([]*Order, error) {
	records, err := OrderRepository.
		Query(&Order{}, nil).
		Where("status IN ?",
			[]string{string(data.OrderStatusQueued), string(data.OrderStatusStockFilled)}).
		Find()

	if err != nil {
		return nil, err
	}

	var orders []*Order
	for _, record := range records {
		orders = append(orders, record.(*Order))
	}

	return orders, nil
}

func (*orderQuery) GetOrderToView(id string) (*Order, error) {
	orderId := data.FromStringToEntityId(id)
	record, err := OrderRepository.
		Query(&Order{}, nil).
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id = ?", orderId).
		First()

	if err != nil {
		return nil, err
	}

	return record.(*Order), nil
}
