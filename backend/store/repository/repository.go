package repository

import (
	"errors"
	data "store/app/domain/data"
	repo "store/app/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresRepository struct{}

func (r *postgresRepository) query(model interface{}, tx repo.Transaction) repo.Query {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	return &query{db: db.Model(model), includeChain: ""}
}

func (r *postgresRepository) create(
	entity data.Entity,
	tx repo.Transaction,
) (string, error) {
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
	id string,
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

func toQueryError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repo.ErrNotFound
	}

	return err
}
