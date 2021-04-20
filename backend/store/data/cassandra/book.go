package cassandra

import (
	"fmt"
	data "store/data"
	"store/utils"
	"store/utils/strings"

	"github.com/scylladb/gocqlx/v2/table"
)

type bookRepository struct {
	cassandraRepository
}

func (r *bookRepository) CreateIfNotExist(
	book data.Book,
	transaction *transaction,
) (data.BookId, error) {
	book.StoreId = data.StoreId()

	_, err := r.cassandraRepository.CreateIfNotExist(book, transaction)
	if err != nil {
		return data.EmptyBookId(), err
	}

	return data.NewBookId(book.GoogleBookId), nil
}

func (r *bookRepository) UpdateOnhandQty(id data.BookId, qty int, transaction *transaction) error {
	command := session.Query(
		r.tableDef.UpdateBuilder().AddLit("onhand_qty", fmt.Sprintf("onhand_qty + %d", qty)).ToCql(),
	).BindMap(
		id.ToMap(),
	)

	return r.cassandraRepository.executeCommand(command, transaction)
}

var bookRepositoryInstance = bookRepository{cassandraRepository{tableDef: table.New(table.Metadata{
	Name: "book",
	Columns: strings.Filter(
		ConvertToColumnNames(utils.FieldsOfType((*data.Book)(nil))),
		func(field string) bool {
			return field != "id"
		},
	),
	PartKey: []string{"store_id"},
	SortKey: []string{"google_book_id"},
})}}
