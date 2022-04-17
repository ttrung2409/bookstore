package messaging

type Message interface {
	Key() string
	Topic() string
}
