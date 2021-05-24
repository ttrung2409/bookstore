package command

type ReceivingBook struct {
	Book
	Qty int
}

type ReceiveBooksRequest struct {
	Items []ReceivingBook
}

type BookReceiving interface {
	Receive(request ReceiveBooksRequest) error
}
