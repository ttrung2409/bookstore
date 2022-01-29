package query

type Query interface {
	FindDeliverableOrders() ([]*Order, error)
	GetOrderDetails(id string) (*Order, error)
}

type query struct{}
