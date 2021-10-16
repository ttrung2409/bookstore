package command

type PlaceAsBackOrderCommand interface {
	Execute(id string) error
}
