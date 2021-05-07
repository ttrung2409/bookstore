package operation

type ReceivingBook struct {
	Book
	Qty int
}

type ReceiveBooksRequest struct {
	Items []ReceivingBook
}

type ReceiveBooks interface {
	Receive(request ReceiveBooksRequest) error
}
