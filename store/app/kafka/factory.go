package kafka

type Factory interface {
	NewProducer(topic string) Producer
	NewConsumer(topic string) Consumer
}
