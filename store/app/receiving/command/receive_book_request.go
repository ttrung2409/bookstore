package command

type ReceiveBooksRequest struct {
	Items []ReceivingBook
}

type ReceivingBook struct {
	Book
	Qty int
}
