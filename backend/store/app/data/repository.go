package data

type repositoryBase interface {
	Get(id EntityId) (interface{}, error)
	Query(entityType interface{}) Query
	Create(entity interface{}, tx Transaction) (EntityId, error)
	Update(id EntityId, entity interface{}, tx Transaction) error
}
