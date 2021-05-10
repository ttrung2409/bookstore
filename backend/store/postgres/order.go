package postgres

import (
	data "store/app/data"
)

type orderRepository struct {
	postgresRepository
}

var orderRepositoryInstance = orderRepository{postgresRepository{newEntity: func() interface{} {
	return &data.Order{}
}}}
