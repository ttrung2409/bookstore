package operation

type AcceptOrder interface {
	Accept(id string) error
}
