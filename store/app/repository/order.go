package repository

import (
	"store/app/domain"
)

type OrderRepository interface {
	Get(id string, tx Transaction) (*domain.Order, error)
	Create(order *domain.Order, tx Transaction) (string, error)
	Update(order *domain.Order, tx Transaction) error
}
