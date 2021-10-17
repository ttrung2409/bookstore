package command

type Command interface {
	Receive(request ReceiveBooksRequest) error
}

type command struct{}
