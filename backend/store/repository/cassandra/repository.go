package repository

import (
	repository "store/repository/interface"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type cassandraRepository struct{}

func (r *cassandraRepository) get(id repository.EntityId, table *table.Table) (interface{}, error) {
	var entity interface{}
	query := session.Query(table.Get()).BindMap(qb.M{"id": id})

	if err := query.GetRelease(&entity); err != nil {
		return nil, err 
	}

	return entity, nil
}

