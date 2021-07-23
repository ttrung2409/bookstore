package command

type AcceptOrderCommand interface {
	Execute(id string) error
}
