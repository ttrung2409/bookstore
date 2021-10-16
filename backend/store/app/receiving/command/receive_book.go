package command

type ReceivingBook struct {
	Book
	Qty int
}

type ReceiveBooksRequest struct {
	Items []ReceivingBook
}

type ReceiveBookCommand interface {
	Execute(request ReceiveBooksRequest) error
}
