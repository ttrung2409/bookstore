package messages

import (
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
