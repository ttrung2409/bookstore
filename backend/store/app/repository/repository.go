package repository

import (
	"store/app/domain/data"
)

type repositoryBase interface {
	create(entity data.Entity, tx Transaction) (string, error)
	update(id string, entity data.Entity, tx Transaction) error
}
