package operation

type OrderCommand interface {
	Accept(id string) error
	PlaceAsBackOrder(id string) error
	Reject(id string) error
}
