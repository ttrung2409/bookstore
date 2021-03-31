package repository

import (
	cassandra "store/repository/cassandra"

	"github.com/sarulabs/di"
)

func Install(builder *di.Builder) {
	cassandra.Install(builder)
}

func Connect() {
	cassandra.Connect();
}