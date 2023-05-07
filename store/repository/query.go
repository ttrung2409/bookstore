package repository

import (
	"fmt"
	"store/app/domain"
	"strings"

	"gorm.io/gorm"
)

type Query[M domain.DataObject] struct {
	db           *gorm.DB
	includeChain string
}

type Where[M domain.DataObject] struct {
	query *Query[M]
	field string
	andOr string
}

func (Query[M]) New() *Query[M] {
	return &Query[M]{db: GetDb().Model(new(M)), includeChain: ""}
}

func (q *Query[M]) Select(fields ...string) *Query[M] {
	q.db = q.db.Select(fields)
	return q
}

func (q *Query[M]) Include(ref string) *Query[M] {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
		q.includeChain = ref
	}

	q.db = q.db.Joins(ref)

	return q
}

func (q *Query[M]) IncludeMany(ref string) *Query[M] {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
		q.includeChain = ref
	}

	q.db = q.db.Preload(ref)

	return q
}

func (q *Query[M]) ThenInclude(relation string) *Query[M] {
	q.includeChain = fmt.Sprintf("%s.%s", q.includeChain, relation)
	return q
}

func (q *Query[M]) Where(field string) *Where[M] {
	return &Where[M]{field: field, query: q, andOr: "and"}
}

func (q *Query[M]) Or(field string) *Where[M] {
	return &Where[M]{field: field, query: q, andOr: "or"}
}

func (q *Query[M]) OrderBy(field string) *Query[M] {
	q.db = q.db.Order(fmt.Sprintf("%s asc", field))
	return q
}

func (q *Query[M]) OrderByDesc(field string) *Query[M] {
	q.db = q.db.Order(fmt.Sprintf("%s desc", field))
	return q
}

func (q *Query[M]) Find() ([]M, error) {
	records, err := q.exec()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (q *Query[M]) First() (M, error) {
	records, err := q.exec()
	if err != nil {
		return *new(M), err
	}

	return records[0], nil
}

func (q *Query[M]) exec() ([]M, error) {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
	}

	var records = []M{}
	if result := q.db.Find(records); result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (c *Where[M]) In(values any) *Query[M] {
	if c.andOr == "and" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s IN ?", c.field), values)
	} else if c.andOr == "or" {
		c.query.db = c.query.db.Or(fmt.Sprintf("%s IN ?", c.field), values)
	}

	return c.query
}

func (c *Where[M]) Eq(value any) *Query[M] {
	if c.andOr == "and" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s = ?", c.field), value)
	} else if c.andOr == "or" {
		c.query.db = c.query.db.Or(fmt.Sprintf("%s = ?", c.field), value)
	}

	return c.query
}

func (c *Where[M]) Contains(value string) *Query[M] {
	if c.andOr == "and" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s LIKE %%?%%", c.field), value)
	} else if c.andOr == "or" {
		c.query.db = c.query.db.Or(fmt.Sprintf("%s LIKE %%?%%", c.field), value)
	}

	return c.query
}
