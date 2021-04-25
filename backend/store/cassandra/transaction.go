package cassandra

import "github.com/gocql/gocql"

type transaction struct {
	commands []Command
}

type transactionFactory struct{}

type Command struct {
	Statement string
	Args      []string
}

func (f *transactionFactory) New() *transaction {
	return &transaction{}
}

func (t *transaction) Commit() error {
	batch := session.NewBatch(gocql.UnloggedBatch)

	for _, command := range t.commands {
		batch.Query(command.Statement, command.Args)
	}

	return session.ExecuteBatch(batch)
}

func (t *transaction) Rollback() {
}
