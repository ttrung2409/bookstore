package repository

import (
	repository "store/repository/interface"

	"store/utils"

	"github.com/gocql/gocql"
	"github.com/sarulabs/di"
	"github.com/scylladb/gocqlx/v2"
)

var session *gocqlx.Session

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*repository.BookRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return &bookRepository{}, nil
			},
		},
		{
			Name: utils.Nameof((*repository.TransactionFactory)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return &transactionFactory{}, nil
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