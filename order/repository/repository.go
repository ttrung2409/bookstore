package repository

import (
	"ecommerce/app/domain"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
)

type cassandraRepository[M domain.DataObject] struct {
	db       *gocqlx.Session
	tableDef *table.Table
}

func (r *cassandraRepository[M]) get(id string) (M, error) {
	var entity M
	query := r.db.Query(r.tableDef.Get()).BindMap(map[string]any{"id": id})

	if err := query.GetRelease(&entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *cassandraRepository) create(
	entity interface{},
	transaction *transaction,
) (string, error) {
	id := data.NewEntityId()

	command := r.session.Query(
		r.tableDef.InsertBuilder().LitColumn("created_at", "toTimestamp(now())").ToCql(),
	).BindStruct(
		entity,
	).BindMap(
		id.ToMap(),
	)

	err := r.executeCommand(command, transaction)

	return id, err
}

func (r *cassandraRepository) createIfNotExist(
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

func (r *cassandraRepository) update(
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
