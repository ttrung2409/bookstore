package data

type repositoryBase interface {
	Get(id EntityId, tx Transaction) (interface{}, error)
	Query(entityType interface{}, tx Transaction) Query
	Create(entity Entity, tx Transaction) (EntityId, error)
	Update(id EntityId, entity Entity, tx Transaction) error
}
