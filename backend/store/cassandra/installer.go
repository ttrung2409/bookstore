package cassandra

import (
	"store/app/data"
	"store/utils"

	"github.com/gocql/gocql"
	"github.com/sarulabs/di"
	"github.com/scylladb/gocqlx/v2"
)

var session *gocqlx.Session

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*data.BookRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return bookRepositoryInstance, nil
			},
		},
		{
			Name: utils.Nameof((*data.BookReceiptRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return bookReceiptRepositoryInstance, nil
			},
		},
		{
			Name: utils.Nameof((*data.OrderRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return orderRepositoryInstance, nil
			},
		},
		{
			Name: utils.Nameof((*data.TransactionFactory)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return transactionFactory{}, nil
			},
		},
	}...)
}

func Connect() *gocqlx.Session {
	if session != nil {
		return session
	}

	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "bookstore"

	sessionValue, err := gocqlx.WrapSession(cluster.CreateSession())

	if err != nil {
		panic(err)
	}

	session = &sessionValue

	return session
}
