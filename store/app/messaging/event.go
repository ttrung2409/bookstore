package messaging

type Event interface {
	Key() string
	Type() string
}
