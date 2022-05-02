package kafka

import (
	"context"
	"log"
	"store/app/kafka"

	kafkaGo "github.com/segmentio/kafka-go"
)

type consumer struct {
	reader *kafkaGo.Reader
}

func newConsumer(topic string) kafka.Consumer {
	return &consumer{reader: kafkaGo.NewReader(kafkaGo.ReaderConfig{
		Brokers:  BrokerAddresses,
		GroupID:  topic,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Logger:   log.Default(),
	})}
}

func (c *consumer) FetchMessage(ctx context.Context) (kafka.Message, error) {
	msg, err := c.reader.FetchMessage(ctx)

	if err != nil {
		return nil, err
	}

	return &message{msg: msg}, nil
}

func (c *consumer) CommitMessage(ctx context.Context, msg kafka.Message) error {
	return c.reader.CommitMessages(ctx, msg.(*message).msg)
}

func (c *consumer) Dispose() error {
	return c.reader.Close()
}
