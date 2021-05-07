package operation

type OrderQuery interface {
	FindByStatus(statuses []string) ([]Order, error)
	Get(id string) (*Order, error)
}
