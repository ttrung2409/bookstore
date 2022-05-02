package messages

import (
	"store/utils"
)

type OrderAccepted struct {
	message
	OrderId string
}

func (e *OrderAccepted) Key() string {
	return e.OrderId
}

func (e *OrderAccepted) Type() string {
	return utils.Nameof(OrderAccepted{})
}
