package repository

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var db *gocqlx.Session

func GetDb() *gocqlx.Session {
	if db == nil {
		db = Connect()
	}

	return db
}

func Connect() *gocqlx.Session {
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "bookstore"

	_db, err := gocqlx.WrapSession(cluster.CreateSession())

	if err != nil {
		panic(err)
	}

	return &_db
}
