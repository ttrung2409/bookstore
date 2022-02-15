package query

import (
	"store/app/domain/data"
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

func (*query) FindDeliverableOrders() ([]*Order, error) {
	var queryFactory = container.Instance().Get(utils.Nameof((*repo.QueryFactory)(nil))).(repo.QueryFactory)

	records, err := queryFactory.
		New(&data.Order{}).
		Where("status").In([]string{string(data.OrderStatusQueued), string(data.OrderStatusStockFilled)}).
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
