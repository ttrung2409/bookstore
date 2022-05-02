package messages

import "store/utils"

type OrderCreated struct {
	message
	OrderId string
}

func (e *OrderCreated) Key() string {
	return e.OrderId
}

func (e *OrderCreated) Type() string {
	return utils.Nameof(OrderCreated{})
}
