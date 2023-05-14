package events

import (
	"store/utils"
)

type OrderDelivered struct {
	OrderId string
}

func (e OrderDelivered) Type() string {
	return utils.Nameof(OrderDelivered{})
}
