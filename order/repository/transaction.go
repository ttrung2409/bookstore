package repository

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Transaction struct {
	db       gocqlx.Session
	commands []struct {
		Statement string
		Args      []string
	}
}

type TransactionalFunc func(tx *Transaction) (any, error)

func (Transaction) New() *Transaction {
	return &Transaction{}
}

func (Transaction) RunInTransaction(
	fn TransactionalFunc,
) (any, error) {
	tx := Transaction{}.New()

	var err error
	defer func() {
		if err != nil {
			tx.rollback()
		}
	}()

	result, err := fn(tx)

	if err != nil {
		return nil, err
	}

	err = tx.commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *Transaction) commit() error {
	batch := t.db.NewBatch(gocql.LoggedBatch)

	for _, command := range t.commands {
		batch.Query(command.Statement, command.Args)
	}

	return t.db.ExecuteBatch(batch)
}
