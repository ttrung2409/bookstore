package repository

import (
	"store/app/domain"
	"store/app/domain/data"
)

type OrderRepository interface {
	repositoryBase
	Get(id data.EntityId, tx Transaction) (*domain.Order, error)
	GetReceivingOrders(tx Transaction) ([]*domain.Order, error)
	Update(order *domain.Order, tx Transaction) error
}
