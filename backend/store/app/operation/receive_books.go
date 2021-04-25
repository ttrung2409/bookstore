package operation

type receivingBook struct {
	Book Book
	Qty  int
}

type ReceiveBooksRequest struct {
	items []receivingBook
}

type ReceiveBooks interface {
	Receive(request ReceiveBooksRequest) error
}
