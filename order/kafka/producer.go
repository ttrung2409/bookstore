package kafka

import (
	"context"
	"log"

	kafkaGo "github.com/segmentio/kafka-go"
	"github.com/thoas/go-funk"
)

type Producer interface {
	Send(ctx context.Context, msgs ...Message) error
	Dispose() error
}

type producer struct {
	writer *kafkaGo.Writer
}

func NewProducer(topic string) Producer {
	return &producer{writer: &kafkaGo.Writer{
		Addr:                   kafkaGo.TCP(ClusterAddress),
		AllowAutoTopicCreation: true,
		Balancer:               &kafkaGo.Hash{},
		Compression:            kafkaGo.Gzip,
		RequiredAcks:           kafkaGo.RequireAll,
		MaxAttempts:            3,
		Topic:                  topic,
		Logger:                 log.Default(),
	}}
}

func (p *producer) Send(ctx context.Context, messages ...Message) error {
	kafkaGoMsgs := funk.Map(messages, func(msg Message) kafkaGo.Message {
		return kafkaGo.Message{
			Key:     []byte(msg.Key()),
			Value:   msg.Value(),
			Headers: []kafkaGo.Header{{Key: "type", Value: []byte(msg.Type())}},
		}
	}).([]kafkaGo.Message)

	return p.writer.WriteMessages(ctx, kafkaGoMsgs...)
}

func (p *producer) Dispose() error {
	return p.writer.Close()
}
