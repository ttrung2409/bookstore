package data

import (
	data "store/data"
	"store/utils"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type bookRepository struct {
	cassandraRepository
}

func (r *bookRepository) GetByGoogleBookId(googleBookId string) (*data.Book, error) {
	var book data.Book
	query := session.Query(r.tableDef.Get()).BindMap(qb.M{"store_id": StoreId(), "google_book_id": googleBookId})

	if err := query.GetRelease(&book); err != nil {
		return nil, err 
	}

	return &book, nil
}

func (r *bookRepository) CreateIfNotExist(book data.Book, transaction *transaction) error {
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