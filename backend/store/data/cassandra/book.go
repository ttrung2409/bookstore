package data

import (
	"fmt"
	data "store/data"
	"store/utils"
	"strings"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type bookRepository struct {
	cassandraRepository
}

func (r *bookRepository) Get(id data.EntityId) (*data.Book, error) {
	var book data.Book

	query := session.Query(r.tableDef.Get()).BindMap(qb.M{"store_id": StoreId(), "google_book_id": googleBookId(id)})

	if err := query.GetRelease(&book); err != nil {
		return nil, err 
	}

	return &book, nil
}

func (r *bookRepository) Create(book data.Book, transaction *transaction) (data.EntityId, error) {
	book.StoreId = StoreId()

	_, err := r.cassandraRepository.Create(book, transaction)

	return r.MakeId(book.GoogleBookId), err
}

func (r *bookRepository) CreateIfNotExist(book data.Book, transaction *transaction) error {
	book.StoreId = StoreId()

	if transaction != nil {
		query := session.Query(r.tableDef.InsertBuilder().Unique().ToCql()).BindStruct(book)
		
		transaction.commands = append(transaction.commands, Command{Statement: query.Statement(), Args: query.Names})

		return nil
	}

	query := session.Query(r.tableDef.InsertBuilder().Unique().ToCql()).BindStruct(book)

	if err := query.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (r *bookRepository) Update(id data.EntityId, book data.Book, transaction *transaction) error {
	return r.cassandraRepository.Update(id, book, transaction)
}

var bookRepositoryInstance = &bookRepository{cassandraRepository{tableDef: table.New(table.Metadata{
	Name: "book", 
	Columns: convertToColumnNames(utils.FieldsOfType((*data.Book)(nil))),
	PartKey: []string{"store_id"},
	SortKey: []string{"google_book_id"},
})}}

func convertToColumnNames(fields []string) []string {
	columns := make([]string, len(fields))

	for _, field := range fields {
		columns = append(columns, FieldNameToColumnName(field))
	}

	return columns
}

func (r *bookRepository) MakeId(googleBookId string) data.EntityId {
	return data.EntityId(fmt.Sprintf("%s@%s", StoreId(), googleBookId))
}

func googleBookId(id data.EntityId) string {
	return strings.Split(id.ToString(), "@")[1]
}