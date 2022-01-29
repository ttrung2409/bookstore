package repository

import (
	"errors"
	data "store/app/domain/data"
	repo "store/app/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresRepository struct {
	newEntity func() data.Entity
}

func (r *postgresRepository) Query(model interface{}, tx repo.Transaction) repo.Query {
	return newQuery(model, tx)
}

func (r *postgresRepository) create(
	entity data.Entity,
	tx repo.Transaction,
) (data.EntityId, error) {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	if result := db.Omit(clause.Associations).Create(entity); result.Error != nil {
		return data.EmptyEntityId, result.Error
	}

	return entity.GetId(), nil
}

func (r *postgresRepository) update(
	id data.EntityId,
	entity data.Entity,
	tx repo.Transaction,
) error {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	entity.SetId(id)
	result := db.Model(entity).Omit(clause.Associations).Updates(entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func toDataQueryError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repo.ErrNotFound
	}

	return err
}
