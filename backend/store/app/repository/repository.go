package repository

import (
	"store/app/domain/data"
)

type repositoryBase interface {
	Query(model interface{}, tx Transaction) Query
	get(id data.EntityId, tx Transaction) (interface{}, error)
	create(entity data.Entity, tx Transaction) (data.EntityId, error)
	update(id data.EntityId, entity data.Entity, tx Transaction) error
}
