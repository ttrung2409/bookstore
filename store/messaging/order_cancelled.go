package messaging

import (
	"store/app/kafka"
	"store/app/order/command"
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
