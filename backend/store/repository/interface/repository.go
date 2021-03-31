package repository

type repository interface {
	Get(id EntityId) (interface{}, error)
	Create(entity interface{}, transaction Transaction) (EntityId, error)
	Update(id EntityId, entity interface{}, transaction Transaction) error
}