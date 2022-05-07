package integration

import (
	"store/app/order/command"
	"store/kafka"
	"store/utils"
)

type OrderCreated struct {
	message
	Order command.Order
}

func (e OrderCreated) Key() string {
	return e.Order.Id
}

func (e OrderCreated) Type() string {
	return utils.Nameof(OrderCreated{})
}

func HandleOrderCreated(msg kafka.Message) error {
	orderCreated := Deserialize(msg, &OrderCreated{})
	command := command.New()

	return command.AcceptOrder(orderCreated.Order)
}
