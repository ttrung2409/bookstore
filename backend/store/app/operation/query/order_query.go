package query

type OrderQuery interface {
	FindOrdersToDeliver() ([]*Order, error)
	GetOrderToView(id string) (*Order, error)
}
