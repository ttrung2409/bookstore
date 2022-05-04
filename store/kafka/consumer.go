package kafka

import (
	"context"
	"log"

	kafkaGo "github.com/segmentio/kafka-go"
)

type Consumer interface {
	FetchMessage(ctx context.Context) (Message, error)
	CommitMessage(ctx context.Context, msg Message) error
	Dispose() error
}

type consumer struct {
	reader *kafkaGo.Reader
}

func NewConsumer(topic string) Consumer {
	return &consumer{reader: kafkaGo.NewReader(kafkaGo.ReaderConfig{
		Brokers:  BrokerAddresses,
		GroupID:  topic,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Logger:   log.Default(),
	})}
}

func (c *consumer) FetchMessage(ctx context.Context) (Message, error) {
	msg, err := c.reader.FetchMessage(ctx)

	if err != nil {
		return nil, err
	}

	return &message{msg: msg}, nil
}

func (c *consumer) CommitMessage(ctx context.Context, msg Message) error {
	return c.reader.CommitMessages(ctx, msg.(*message).msg)
}

func (c *consumer) Dispose() error {
	return c.reader.Close()
}
