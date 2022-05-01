package repository

import (
	"fmt"
	repo "store/app/repository"
	"strings"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type query struct {
	db           *gorm.DB
	includeChain string
}

type queryCondition struct {
	query *query
	field string
	andOr string
}

func (q *query) Select(fields ...string) repo.Query {
	q.db = q.db.Select(fields)
	return q
}

func (q *query) Include(ref string) repo.Query {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
		q.includeChain = ref
	}

	q.db = q.db.Joins(ref)

	return q
}

func (q *query) IncludeMany(ref string) repo.Query {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
		q.includeChain = ref
	}

	q.db = q.db.Preload(ref)

	return q
}

func (q *query) ThenInclude(relation string) repo.Query {
	q.includeChain = fmt.Sprintf("%s.%s", q.includeChain, relation)
	return q
}

func (q *query) Where(field string) repo.QueryCondition {
	return &queryCondition{field: field, query: q, andOr: "and"}
}

func (q *query) Or(field string) repo.QueryCondition {
	return &queryCondition{field: field, query: q, andOr: "or"}
}

func (q *query) OrderBy(field string) repo.Query {
	q.db = q.db.Order(fmt.Sprintf("%s asc", field))
	return q
}

func (q *query) OrderByDesc(field string) repo.Query {
	q.db = q.db.Order(fmt.Sprintf("%s desc", field))
	return q
}

func (q *query) Find() ([]interface{}, error) {
	records, err := q.exec()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (q *query) First() (interface{}, error) {
	records, err := q.exec()
	if err != nil {
		return nil, err
	}

	return records[0], nil
}

func (q *query) exec() ([]interface{}, error) {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
	}

	var records []interface{}
	if result := q.db.Find(records); result.Error != nil {
		return nil, toQueryError(result.Error)
	}

	return funk.Map(records, func(record interface{}) interface{} {
		return &record
	}).([]interface{}), nil
}

func (c *queryCondition) In(values interface{}) repo.Query {
	if c.andOr == "and" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s IN ?", c.field), values)
	} else if c.andOr == "or" {
		c.query.db = c.query.db.Or(fmt.Sprintf("%s IN ?", c.field), values)
	}

	return c.query
}

func (c *queryCondition) Eq(value interface{}) repo.Query {
	if c.andOr == "and" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s = ?", c.field), value)
	} else if c.andOr == "or" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s = ?", c.field), value)
	}

	return c.query
}

func (c *queryCondition) Contains(value string) repo.Query {
	if c.andOr == "and" {
		c.query.db = c.query.db.Where(fmt.Sprintf("%s LIKE %%?%%", c.field), value)
	} else if c.andOr == "or" {
		c.query.db = c.query.db.Or(fmt.Sprintf("%s LIKE %%?%%", c.field), value)
	}

	return c.query
}

type queryFactory struct{}

func (*queryFactory) New(model interface{}) repo.Query {
	return &query{db: db.Model(model), includeChain: ""}
}
