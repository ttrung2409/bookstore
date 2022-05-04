package integration

import (
	"fmt"
	"store/kafka"
	"store/utils"
)

type MessageHandler func(msg kafka.Message) error

func NewHandler(msg kafka.Message) (MessageHandler, error) {
	switch msg.Type() {
	case utils.Nameof(OrderCreated{}):
		return HandleOrderCreated, nil
	case utils.Nameof(OrderCancelled{}):
		return HandleOrderCancelled, nil
	default:
		return nil, fmt.Errorf("no handler found for message type: %s", msg.Type())
	}
}
