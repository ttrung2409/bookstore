package query

type OrderQuery interface {
	FindOrdersToDeliver() ([]*Order, error)
	GetOrderDetails(id string) (*Order, error)
}
