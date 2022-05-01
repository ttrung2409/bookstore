package container

import (
	"github.com/sarulabs/di"
)

var container di.Container
var builder *di.Builder

func ContainerBuilder() *di.Builder {
	if builder != nil {
		return builder
	}

	builder, _ = di.NewBuilder()

	return builder
}

func Instance() di.Container {
	if container != nil {
		return container
	}

	return ContainerBuilder().Build()
}
