package messaging

type Producer interface {
	Send(messages ...interface{}) error
}
