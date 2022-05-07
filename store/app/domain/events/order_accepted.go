package events

import (
	"store/utils"
)

type OrderAccepted struct {
	OrderId string
}

func (e OrderAccepted) Type() string {
	return utils.Nameof(OrderAccepted{})
}
