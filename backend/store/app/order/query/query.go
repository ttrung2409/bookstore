package query

import (
	"store/app/domain/data"
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

type Query interface {
	FindDeliverableOrders() ([]*Order, error)
	GetOrderDetails(id string) (*Order, error)
}

func New() Query {
	return &query{}
}

type query struct{}

func (*query) FindDeliverableOrders() ([]*Order, error) {
	queryFactory := container.Instance().Get(utils.Nameof((*repo.QueryFactory)(nil))).(repo.QueryFactory)

	records, err := queryFactory.
		New(&data.Order{}).
		Where("status").Eq(data.OrderStatusAccepted).
		Find()

	if err != nil {
		return nil, err
	}

	var orders []*Order
	for _, record := range records {
		order := Order{}.fromDataObject(record.(*data.Order))
		orders = append(orders, order)
	}

	return orders, nil
}

func (*query) GetOrderDetails(id string) (*Order, error) {
	queryFactory := container.Instance().Get(utils.Nameof((*repo.QueryFactory)(nil))).(repo.QueryFactory)

	record, err := queryFactory.
		New(&data.Order{}).
		Include("Customer").
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id").Eq(id).
		First()

	if err != nil {
		return nil, err
	}

	order := Order{}.fromDataObject(record.(*data.Order))

	return order, nil
}
