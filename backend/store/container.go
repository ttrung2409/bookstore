package container

import (
	"store/repository"

	"github.com/sarulabs/di"
)

var Container di.Container

func install() {
	builder, _ := di.NewBuilder()
	repository.Install(builder)
	Container = builder.Build()

	repository.Connect()
}
