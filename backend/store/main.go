package main

import (
	order "store/app/order"
	receiving "store/app/receiving"
	"store/container"
	postgres "store/repository/postgres"
)

func main() {
	builder := container.ContainerBuilder()

	postgres.Install(builder)
	receiving.Install(builder)
	order.Install(builder)
}
