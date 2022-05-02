package kafka

import "store/app/kafka"

type factory struct{}

func (*factory) NewProducer(topic string) kafka.Producer {
	return newProducer(topic)
}

func (*factory) NewConsumer(topic string) kafka.Consumer {
	return newConsumer(topic)
}
