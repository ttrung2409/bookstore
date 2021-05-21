package data

type repositoryBase interface {
	Query(entityType interface{}, tx Transaction) Query
	get(id EntityId, tx Transaction) (interface{}, error)
	create(entity Entity, tx Transaction) (EntityId, error)
	update(id EntityId, entity Entity, tx Transaction) error
}
