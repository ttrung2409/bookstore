package repository

import (
	"fmt"
	"store/app/domain/data"
	repo "store/app/repository"
	"strings"

	"gorm.io/gorm"
)

type query[M data.Model] struct {
	db           *gorm.DB
	includeChain string
}

type where[M data.Model] struct {
	query *query[M]
	field string
	andOr string
}

func (q *query[M]) Select(fields ...string) repo.Query[M] {
	q.db = q.db.Select(fields)
	return q
}

func (q *query[M]) Include(ref string) repo.Query[M] {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
		q.includeChain = ref
	}

	q.db = q.db.Joins(ref)

	return q
}

func (q *query[M]) IncludeMany(ref string) repo.Query[M] {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
		q.includeChain = ref
	}

	q.db = q.db.Preload(ref)

	return q
}

func (q *query[M]) ThenInclude(relation string) repo.Query[M] {
	q.includeChain = fmt.Sprintf("%s.%s", q.includeChain, relation)
	return q
}

func (q *query[M]) Where(field string) repo.Where[M] {
	return &where[M]{field: field, query: q, andOr: "and"}
}

func (q *query[M]) Or(field string) repo.Where[M] {
	return &where[M]{field: field, query: q, andOr: "or"}
}

func (q *query[M]) OrderBy(field string) repo.Query[M] {
	q.db = q.db.Order(fmt.Sprintf("%s asc", field))
	return q
}

func (q *query[M]) OrderByDesc(field string) repo.Query[M] {
	q.db = q.db.Order(fmt.Sprintf("%s desc", field))
	return q
}

func (q *query[M]) Find() ([]M, error) {
	records, err := q.exec()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (q *query[M]) First() (M, error) {
	records, err := q.exec()
	if err != nil {
		return *new(M), err
	}

	return records[0], nil
}

func (q *query[M]) exec() ([]M, error) {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
	}

	var records = []M{}
	if result := q.db.Find(records); result.Error != nil {
		return nil, toQueryError(result.Error)
	}

	return records, nil
}

func (c *where[M]) In(values interface{}) repo.Query[M] {
	if c.andOr == "and" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s IN ?", c.field), values)
	} else if c.andOr == "or" {
		c.query.db = c.query.db.Or(fmt.Sprintf("%s IN ?", c.field), values)
	}

	return c.query
}

func (c *where[M]) Eq(value interface{}) repo.Query[M] {
	if c.andOr == "and" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s = ?", c.field), value)
	} else if c.andOr == "or" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s = ?", c.field), value)
	}

	return c.query
}

func (c *where[M]) Contains(value string) repo.Query[M] {
	if c.andOr == "and" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s LIKE %%?%%", c.field), value)
	} else if c.andOr == "or" {
		c.query.db = c.query.db.Or(fmt.Sprintf("%s LIKE %%?%%", c.field), value)
	}

	return c.query
}

type queryFactory[M data.Model] struct{}

func (queryFactory[M]) New() repo.Query[M] {
	return &query[M]{db: Db().Model(new(M)), includeChain: ""}
}
