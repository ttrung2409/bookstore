package event

import (
	"fmt"
	"store/app/domain/data"
)

type OrderStatusChanged struct {
	OrderId string
	Status  data.OrderStatus
}

func (event OrderStatusChanged) Key() string {
	return fmt.Sprintf("%s-%s", event.OrderId, event.Status)
}

func (event OrderStatusChanged) Topic() string {
	return "order"
}
