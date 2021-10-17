package command

type Command interface {
	AcceptOrder(orderId string) error
	PlaceAsBackOrder(orderId string) error
	RejectOrder(orderId string) error
}

type command struct{}
