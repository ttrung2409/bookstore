package events

import (
	"store/utils"
)

type OrderRejected struct {
	OrderId string
}

func (e *OrderRejected) Key() string {
	return e.OrderId
}

func (e *OrderRejected) Type() string {
	return utils.Nameof(OrderRejected{})
}
