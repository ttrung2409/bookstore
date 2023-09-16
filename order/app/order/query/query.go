package query

type Query interface {
	GetOrderDetails(id string) error
}
