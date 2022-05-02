package kafka

type Message interface {
	Key() string
	Type() string
	Value() []byte
}
