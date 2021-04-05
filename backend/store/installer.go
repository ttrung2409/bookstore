package container

import (
	data "store/data/cassandra"

	"github.com/sarulabs/di"
)

var Container di.Container

func install() {
	builder, _ := di.NewBuilder()
	data.Install(builder)
	Container = builder.Build()

	data.Connect()
}
