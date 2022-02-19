package repository

import (
	"store/app/domain"
)

type OrderRepository interface {
	repositoryBase
	Get(id string, tx Transaction) (*domain.Order, error)
	GetReceivingOrders(tx Transaction) ([]*domain.Order, error)
	Update(order *domain.Order, tx Transaction) error
}
