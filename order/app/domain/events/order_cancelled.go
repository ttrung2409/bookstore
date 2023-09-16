package events

import "ecommerce/utils"

type OrderCancelled struct {
	OrderId string
}

func (e OrderCancelled) Type() string {
	return utils.Nameof(OrderCancelled{})
}
