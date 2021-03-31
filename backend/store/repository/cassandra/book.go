package repository

import (
	repository "store/repository/interface"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type bookRepository struct {
	cassandraRepository
}

var schema = table.New(table.Metadata{
	Name: "book", 
	Columns: []string{
		"store_id", 
		"google_book_id",
		"authors", 
		"average_rating", 
		"description",
		"preserved_qty",
		"preview_url",
		"published_date",
		"publisher",
		"ratings_count",
		"subtitle",
		"thumbnail_url",
		"title",
	},
	PartKey: []string{"store_id"},
	SortKey: []string{"google_book_id"},
})

func (r *bookRepository) GetByGoogleBookId(googleBookId string) (interface{}, error) {
	var entity interface{}
	query := session.Query(schema.Get()).BindMap(qb.M{"store_id": StoreId(), "google_book_id": googleBookId})

	if err := query.GetRelease(&entity); err != nil {
		return nil, err 
	}

	return entity, nil
}

func (r *bookRepository) Create(book interface{}, transaction *transaction) (repository.EntityId, error) {
	if transaction != nil {
		query := session.Query(schema.Insert()).BindStruct(book)
		
		transaction.commands = append(transaction.commands, Command{Statement: query.Statement(), Args: query.Names})

		return entityId, nil
	}

	query := session.Query(schema.Insert()).BindStruct(book)

	if err := query.ExecRelease(); err != nil {
		return repository.EmptyEntityId, err
	}

	return entityId, nil
}

func (r *bookRepository) Update(book interface{}, transaction *transaction) (repository.EntityId, error) {
	
}