package repository

import (
	"errors"
	"store/app/domain/data"
	repo "store/app/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresRepository[M data.Model] struct{}

func (r *postgresRepository[M]) query(tx repo.Transaction) repo.Query[M] {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	return &query[M]{db: db.Model(new(M)), includeChain: ""}
}

func (r *postgresRepository[M]) create(
	entity M,
	tx repo.Transaction,
) error {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	if result := db.Omit(clause.Associations).Create(&entity); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgresRepository[M]) update(
	entity M,
	tx repo.Transaction,
) error {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	result := db.Model(new(M)).Omit(clause.Associations).Updates(entity)
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
