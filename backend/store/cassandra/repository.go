package cassandra

import (
	"errors"
	data "store/app/data"
	"store/utils"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
)

type cassandraRepository struct {
	tableDef *table.Table
}

func (r *cassandraRepository) Get(id data.Identifier) (interface{}, error) {
	var entity interface{}
	query := session.Query(r.tableDef.Get()).BindMap(id.ToMap())

	if err := query.GetRelease(&entity); err != nil {
		return nil, r.convertToDataError(err)
	}

	return entity, nil
}

func (r *cassandraRepository) Create(
	entity interface{},
	transaction *transaction,
) (data.Identifier, error) {
	id := data.NewEntityId()

	command := session.Query(
		r.tableDef.InsertBuilder().LitColumn("created_at", "toTimestamp(now())").ToCql(),
	).BindStruct(
		entity,
	).BindMap(
		id.ToMap(),
	)

	err := r.executeCommand(command, transaction)

	return id, err
}

func (r *cassandraRepository) CreateIfNotExist(
	entity interface{},
	transaction *transaction,
) (data.Identifier, error) {
	id := data.NewEntityId()

	command := session.Query(
		r.tableDef.InsertBuilder().LitColumn("created_at", "toTimestamp(now())").Unique().ToCql(),
	).BindStruct(
		entity,
	).BindMap(
		id.ToMap(),
	)

	err := r.executeCommand(command, transaction)

	return id, err
}

func (r *cassandraRepository) Update(
	id data.Identifier,
	entity interface{},
	transaction *transaction,
) error {
	command := session.Query(
		r.tableDef.Update(utils.FieldsOfObject(entity)...),
	).BindStruct(
		entity,
	).BindMap(
		id.ToMap(),
	)

	err := r.executeCommand(command, transaction)

	return err
}

func (r cassandraRepository) executeCommand(
	command *gocqlx.Queryx,
	transaction *transaction,
) error {
	if transaction != nil {
		transaction.commands = append(
			transaction.commands,
			Command{Statement: command.Statement(), Args: command.Names},
		)

		return nil
	}

	return command.ExecRelease()
}

func (r cassandraRepository) convertToDataError(err error) error {
	if errors.Is(err, gocql.ErrNotFound) {
		return data.ErrNotFound
	}

	return err
}
