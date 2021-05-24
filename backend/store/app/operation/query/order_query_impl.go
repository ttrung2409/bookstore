package query

import (
	"store/app/data"
)

type orderQuery struct{}

func (*orderQuery) FindOrdersToDeliver() ([]*Order, error) {
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
		orders = append(orders, Order{}.fromDataObject(dataOrder))
	}

	return orders, nil
}

func (*orderQuery) GetOrderToView(id string) (*Order, error) {
	orderId := data.FromStringToEntityId(id)
	record, err := OrderRepository.
		Query(&data.Order{}, nil).
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id = ?", orderId).
		First()

	if err != nil {
		return nil, err
	}

	order := Order{}.fromDataObject(record.(*data.Order))

	return order, nil
}
