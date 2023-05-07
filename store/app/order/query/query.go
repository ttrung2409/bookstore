package query

import (
	"store/app/domain"
	repo "store/repository"
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
	records, err := repo.Query[domain.OrderData]{}.New().
		Where("status = ?", domain.OrderStatusAccepted).
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
	record, err := repo.Query[domain.OrderData]{}.New().
		Join("Customer").
		Preload("Items").
		Join("Items.Book").
		Where("id = ?", id).
		FindOne()

	if err != nil {
		return Order{}, err
	}

	order := Order{}.fromDataObject(record)

	return order, nil
}
