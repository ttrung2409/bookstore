package query

type Query interface {
	FindOrdersToDeliver() ([]*Order, error)
	GetOrderDetails(id string) (*Order, error)
}

type query struct{}
