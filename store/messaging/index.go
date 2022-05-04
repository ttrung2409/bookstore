package messaging

import (
	"context"
	"fmt"
	"store/app/kafka"
	"store/container"
	"store/utils"
)

var consumers map[string]kafka.Consumer

func Start(ctx context.Context) {
	if consumers == nil {
		consumers = map[string]kafka.Consumer{}
	}

	startConsumer(ctx, "order")
}

func Stop() {
	if consumers == nil {
		return
	}

	for _, consumer := range consumers {
		consumer.Dispose()
	}

	consumers = nil
}

func startConsumer(ctx context.Context, topic string) error {
	if _, ok := consumers[topic]; ok {
		return fmt.Errorf("consumer of topic %s has already been started\n", topic)
	}

	factory := container.Instance().Get(utils.Nameof((*kafka.Factory)(nil))).(kafka.Factory)
	consumers[topic] = factory.NewConsumer(topic)

	for {
		msg, err := consumers[topic].FetchMessage(ctx)
		if err != nil {
			return err
		}

		handler, err := NewHandler(msg)
		if err != nil {
			return err
		}

		if err := handler(msg); err != nil {
			return err
		}

		if err := consumers[topic].CommitMessage(ctx, msg); err != nil {
			return err
		}
	}
}
