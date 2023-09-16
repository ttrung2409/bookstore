package events

import (
	"ecommerce/app/domain"
	"ecommerce/utils"
)

type OrderCreated struct {
	Order domain.BookData
}

func (e OrderCreated) Type() string {
	return utils.Nameof(OrderCreated{})
}
