package repository

import (
	"ecommerce/app/domain"
	"ecommerce/utils"

	"github.com/scylladb/gocqlx/v2/table"
)

type OrderRepository interface {
	Get(id string, tx Transaction) (*domain.Order, error)
	Create(order *domain.Order, tx Transaction) error
}

type orderRepository struct {
	cassandraRepository
}

var orderRepositoryInstance = orderRepository{cassandraRepository{tableDef: table.New(table.Metadata{
	Name:    "order",
	Columns: ConvertToColumnNames(utils.FieldsOfType((*domain.OrderData)(nil))),
	PartKey: []string{"customer_id"},
	SortKey: []string{"created_at"},
})}}
