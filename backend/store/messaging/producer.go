package messaging

import (
	"context"
	"encoding/json"
	"store/app/messaging"

	"github.com/segmentio/kafka-go"
	"github.com/thoas/go-funk"
)

type producer struct {
}

func createWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.Hash{},
		Compression:            kafka.Snappy,
		RequiredAcks:           kafka.RequireOne,
		MaxAttempts:            3,
	}
}

func (*producer) Send(messages ...messaging.Message) (err error) {
	writer := createWriter()

	defer func() {
		closeErr := writer.Close()
		if err == nil {
			err = closeErr
		}
	}()

	kafkaMsgs := funk.Map(messages, func(msg messaging.Message) kafka.Message {
		value, _ := json.Marshal(msg)
		return kafka.Message{
			Topic: msg.Topic(),
			Key:   []byte(msg.Key()),
			Value: value,
		}
	}).([]kafka.Message)

	err = writer.WriteMessages(context.Background(), kafkaMsgs...)

	return err
}
