package query

type OrderQuery interface {
	FindByStatus(statuses []string) ([]Order, error)
	GetWithItems(id string) (*Order, error)
}
