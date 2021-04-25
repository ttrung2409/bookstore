package container

import (
	cassandra "store/cassandra"

	"github.com/sarulabs/di"
)

var Container di.Container

func main() {
	builder, _ := di.NewBuilder()
	cassandra.Install(builder)
	Container = builder.Build()

	cassandra.Connect()
}
