package cassandra

import (
	data "store/app/data"
	"store/utils"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type orderRepository struct {
	cassandraRepository
}

func (r *orderRepository) FindByStatus(statuses []string) ([]data.Order, error) {
	var orders []data.Order
	query := session.Query(
		r.tableDef.SelectBuilder().Where(qb.In("status")).ToCql(),
	).BindMap(
		qb.M{"status": statuses},
	)

	if err := query.GetRelease(&orders); err != nil {
		return nil, r.convertToDataError(err)
	}

	return orders, nil
}

var orderRepositoryInstance = orderRepository{cassandraRepository{tableDef: table.New(table.Metadata{
	Name:    "order",
	Columns: ConvertToColumnNames(utils.FieldsOfType((*data.Order)(nil))),
	PartKey: []string{"status"},
	SortKey: []string{"id"},
})}}
