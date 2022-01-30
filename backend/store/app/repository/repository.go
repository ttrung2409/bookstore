package repository

import (
	"store/app/domain/data"
)

type repositoryBase interface {
	create(entity data.Entity, tx Transaction) (data.EntityId, error)
	update(id data.EntityId, entity data.Entity, tx Transaction) error
}
