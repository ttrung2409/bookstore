package kafka

import (
	kafkaGo "github.com/segmentio/kafka-go"
	"github.com/thoas/go-funk"
)

type Message interface {
	Key() string
	Type() string
	Value() []byte
}

type message struct {
	msg kafkaGo.Message
}

func (m *message) Key() string {
	return string(m.msg.Key)
}

func (m *message) Type() string {
	header := funk.Find(m.msg.Headers, func(header kafkaGo.Header) bool {
		return header.Key == "type"
	}).(kafkaGo.Header)

	return string(header.Value)
}

func (m *message) Value() []byte {
	return m.msg.Value
}
