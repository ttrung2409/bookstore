package data

import (
	data "store/data"
	"store/utils"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type cassandraRepository struct {
	tableDef *table.Table
}

func (r *cassandraRepository) Get(id data.EntityId) (interface{}, error) {
	var entity interface{}
	query := session.Query(r.tableDef.Get()).BindMap(qb.M{"id": id})

	if err := query.GetRelease(&entity); err != nil {
		return nil, err 
	}

	return entity, nil
}

func (r *cassandraRepository) Create(entity interface{}, transaction *transaction) (data.EntityId, error) {
	id := data.NewEntityId()

	if transaction != nil {
		query := session.Query(r.tableDef.Insert()).BindStruct(entity).BindMap(id.ToMap())
		
		transaction.commands = append(transaction.commands, Command{Statement: query.Statement(), Args: query.Names})

		return id, nil
	}

	query := session.Query(r.tableDef.Insert()).BindStruct(entity).BindMap(id.ToMap())

	if err := query.ExecRelease(); err != nil {
		return id, err
	}

	return id, nil
}


func (r *cassandraRepository) Update(id data.EntityId, entity interface{}, transaction *transaction) error {
	if transaction != nil {
		query := session.Query(r.tableDef.Update(utils.FieldsOfObject(entity)...)).BindStruct(entity).BindMap(id.ToMap())
		
		transaction.commands = append(transaction.commands, Command{Statement: query.Statement(), Args: query.Names})

		return nil
	}

	query := session.Query(r.tableDef.Update(utils.FieldsOfObject(entity)...)).BindStruct(entity).BindMap(qb.M{"id": id})

	if err := query.ExecRelease(); err != nil {
		return err
	}

	return nil
}

