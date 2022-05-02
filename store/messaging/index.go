package messaging

import (
	"context"
	"encoding/json"
	"store/app/kafka"
	"store/app/order/command"
	"store/container"
	"store/messaging/messages"
	"store/utils"
)

var consumer kafka.Consumer

func Start(ctx context.Context) {
	startConsumer(ctx, "order")
}

func Stop() {
	if consumer == nil {
		return
	}

	consumer.Dispose()
	consumer = nil
}

func startConsumer(ctx context.Context, topic string) error {
	factory := container.Instance().Get(utils.Nameof((*kafka.Factory)(nil))).(kafka.Factory)
	consumer = factory.NewConsumer(topic)

	for {
		msg, err := consumer.FetchMessage(ctx)
		if err != nil {
			return err
		}

		switch msg.Type() {
		case utils.Nameof(messages.OrderCancelled{}):
			command := command.New()

			orderCancelled := &messages.OrderCancelled{}
			json.Unmarshal(msg.Value(), orderCancelled)

			if err := command.CancelOrder(orderCancelled.OrderId); err == nil {
				consumer.CommitMessage(ctx, msg)
			}
		}
	}
}
