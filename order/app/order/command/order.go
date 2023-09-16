package command

type Order struct {
	Customer Customer
	Items    []OrderItem
}

type OrderItem struct {
	BookId string
	Qty    int
}
