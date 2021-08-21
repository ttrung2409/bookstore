package container

import (
	postgres "store/repository/postgres"

	"github.com/sarulabs/di"
)

var container di.Container

func main() {
	builder, _ := di.NewBuilder()
	postgres.Install(builder)
	container = builder.Build()
}

func Container() di.Container {
	return container
}
