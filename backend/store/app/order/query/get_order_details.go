package query

import (
	"store/app/domain/data"
)

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
