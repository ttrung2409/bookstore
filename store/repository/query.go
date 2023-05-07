package repository

import (
	"store/app/domain"

	"gorm.io/gorm"
)

type Query[M domain.DataObject] struct {
	db *gorm.DB
}

func (Query[M]) New() *Query[M] {
	return &Query[M]{GetDb().Model(new(M))}
}

func (q *Query[M]) Select(fields ...string) *Query[M] {
	q.db = q.db.Select(fields)
	return q
}

func (q *Query[M]) Preload(relation string, args ...any) *Query[M] {
	q.db = q.db.Preload(relation, args...)
	return q
}

func (q *Query[M]) Join(relation string, args ...any) *Query[M] {
	q.db = q.db.Joins(relation, args...)
	return q
}

func (q *Query[M]) Where(filters any, args ...any) *Query[M] {
	q.db = q.db.Where(filters, args...)
	return q
}

func (q *Query[M]) Or(filters any, args ...any) *Query[M] {
	q.db = q.db.Or(filters, args...)
	return q
}

func (q *Query[M]) Order(value any) *Query[M] {
	q.db = q.db.Order(value)
	return q
}

func (q *Query[M]) Find() ([]M, error) {
	var records = []M{}
	if result := q.db.Find(records); result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (q *Query[M]) FindOne() (M, error) {
	var records = []M{}
	if result := q.db.Find(records); result.Error != nil {
		return *new(M), result.Error
	}

	return records[0], nil
}
