package messaging

type Producer interface {
	Send(messages ...Message) error
}
