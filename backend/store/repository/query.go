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

func (q *query) Where(condition string, args ...interface{}) repo.Query {
	q.db = q.db.Where(condition, args)
	return q
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
		return nil, toDataQueryError(result.Error)
	}

	return funk.Map(records, func(record interface{}) interface{} {
		return &record
	}).([]interface{}), nil
}

type queryFactory struct{}

func (*queryFactory) New(model interface{}) repo.Query {
	return &query{db: db.Model(model), includeChain: ""}
}
