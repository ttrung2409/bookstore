package event

import (
	"store/app/domain/data"
	"store/app/messaging"
)

type OrderStatusChanged struct {
	*messaging.Message
	OrderId string
	Status  data.OrderStatus
}
