package data

type repositoryBase interface {
	Get(id EntityId) (interface{}, error)
	Query(entityType interface{}) Query
	Create(entity Entity, tx Transaction) (EntityId, error)
	Update(id EntityId, entity Entity, tx Transaction) error
}
