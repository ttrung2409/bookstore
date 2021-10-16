package command

type RejectOrderCommand interface {
	Execute(id string) error
}
