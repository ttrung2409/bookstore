package repository

type Repository interface {
	Get(id string) struct{}
	Create(entity struct{}) string
	Update(id string, entity struct{})
}