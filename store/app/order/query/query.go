package query

import (
	"store/app/domain/data"
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

type Query interface {
	FindDeliverableOrders() ([]Order, error)
	GetOrderDetails(id string) (Order, error)
}

func New() Query {
	return &query{}
}

type query struct{}

func (*query) FindDeliverableOrders() ([]Order, error) {
	queryFactory := container.Instance().Get(utils.Nameof((*repo.QueryFactory[data.Order])(nil))).(repo.QueryFactory[data.Order])

	records, err := queryFactory.
		New().
		Where("status").Eq(data.OrderStatusAccepted).
		Find()

	if err != nil {
		return nil, err
	}

	var orders []Order
	for _, record := range records {
		order := Order{}.fromDataObject(record)
		orders = append(orders, order)
	}

	return orders, nil
}

func (*query) GetOrderDetails(id string) (Order, error) {
	queryFactory := container.Instance().Get(utils.Nameof((*repo.QueryFactory[data.Order])(nil))).(repo.QueryFactory[data.Order])

	record, err := queryFactory.
		New().
		Include("Customer").
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id").Eq(id).
		First()

	if err != nil {
		return Order{}, err
	}

	order := Order{}.fromDataObject(record)

	return order, nil
}
