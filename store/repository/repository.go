package repository

import (
	"store/app/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresRepository[M domain.DataObject] struct {
	db *gorm.DB
}

func (r *postgresRepository[M]) query(tx *Transaction) *Query[M] {
	db := r.db
	if tx != nil {
		db = tx.db
	}

	return &Query[M]{db.Model(new(M))}
}

func (r *postgresRepository[M]) create(
	entity M,
	tx *Transaction,
) error {
	db := r.db
	if tx != nil {
		db = tx.db
	}

	if result := db.Omit(clause.Associations).Create(&entity); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgresRepository[M]) update(
	entity M,
	tx *Transaction,
) error {
	db := r.db
	if tx != nil {
		db = tx.db
	}

	if result := db.Model(new(M)).Omit(clause.Associations).Updates(entity); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *postgresRepository[M]) batchDelete(tx *Transaction, where string, args ...any) error {
	db := r.db
	if tx != nil {
		db = tx.db
	}

	if result := db.Where(where, args...).Delete(new(M)); result.Error != nil {
		return result.Error
	}

	return nil
}
