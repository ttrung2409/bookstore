package postgres

import (
	"fmt"
	"store/app/data"
	"strings"

	"gorm.io/gorm"
)

type query struct {
	db           *gorm.DB
	includeChain string
}

func newQuery(model interface{}, tx data.Transaction) data.Query {
	db := Db()
	if tx != nil {
		db = tx.(*transaction).db
	}

	return &query{db: db.Model(model), includeChain: ""}
}

func (q *query) Select(fields ...string) data.Query {
	q.db = q.db.Select(fields)
	return q
}

func (q *query) Include(ref string) data.Query {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
		q.includeChain = ref
	}

	q.db = q.db.Joins(ref)

	return q
}

func (q *query) IncludeMany(ref string) data.Query {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
		q.includeChain = ref
	}

	q.db = q.db.Preload(ref)

	return q
}

func (q *query) ThenInclude(relation string) data.Query {
	q.includeChain = fmt.Sprintf("%s.%s", q.includeChain, relation)
	return q
}

func (q *query) Where(condition string, args ...interface{}) data.Query {
	q.db = q.db.Where(condition, args)
	return q
}

func (q *query) OrderBy(field string) data.Query {
	q.db = q.db.Order(fmt.Sprintf("%s asc", field))
	return q
}

func (q *query) OrderByDesc(field string) data.Query {
	q.db = q.db.Order(fmt.Sprintf("%s desc", field))
	return q
}

func (q *query) Find() ([]interface{}, error) {
	result, err := q.exec()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (q *query) First() (interface{}, error) {
	result, err := q.exec()
	if err != nil {
		return nil, err
	}

	return result[0], nil
}

func (q *query) exec() ([]interface{}, error) {
	if q.includeChain != "" && strings.Contains(q.includeChain, ".") {
		q.db = q.db.Preload(q.includeChain)
	}

	var records []interface{}
	if result := q.db.Find(records); result.Error != nil {
		return nil, toDataQueryError(result.Error)
	}

	return records, nil
}
