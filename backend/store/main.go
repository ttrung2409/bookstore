package main

import (
	order "store/app/order"
	receiving "store/app/receiving"
	"store/container"
	repository "store/repository"
)

func main() {
	builder := container.ContainerBuilder()

	repository.Install(builder)
	receiving.Install(builder)
	order.Install(builder)
}
