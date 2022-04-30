package main

import (
	"store/container"
	repository "store/repository"
	server "store/rest"
)

func main() {
	builder := container.ContainerBuilder()

	repository.Install(builder)

	server.Start()
}
