package query

import (
	"store/app/domain/data"
)

func (*query) FindOrdersToDeliver() ([]*Order, error) {
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

func (*query) GetOrderDetails(id string) (*Order, error) {
	orderId := data.FromStringToEntityId(id)
	record, err := OrderRepository.
		Query(&data.Order{}, nil).
		Include("Customer").
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id = ?", orderId).
		First()

	if err != nil {
		return nil, err
	}

	order := Order{}.fromDataObject(record.(*data.Order))

	return &order, nil
}
