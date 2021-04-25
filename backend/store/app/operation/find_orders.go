package operation

type FindOrders interface {
	Find(status string) ([]Order, error)
}
