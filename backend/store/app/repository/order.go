package repository

import "store/app/domain/data"

type OrderRepository interface {
	repositoryBase
	Get(id data.EntityId, tx Transaction) (*data.Order, error)
	GetReceivingOrders(tx Transaction) ([]*data.Order, error)
	Update(order *data.Order, tx Transaction) error
}
