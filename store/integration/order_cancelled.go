package integration

import (
	"store/app/order/command"
	"store/kafka"
	"store/utils"
)

type OrderCancelled struct {
	message
	OrderId string
}

func (e *OrderCancelled) Key() string {
	return e.OrderId
}

func (e *OrderCancelled) Type() string {
	return utils.Nameof(OrderCancelled{})
}

func HandleOrderCancelled(msg kafka.Message) error {
	orderCancelled := Deserialize(msg, &OrderCancelled{})
	command := command.New()

	return command.CancelOrder(orderCancelled.OrderId)
}
