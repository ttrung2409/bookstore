package cassandra

import (
	data "store/app/data"
	"store/utils"

	"github.com/scylladb/gocqlx/v2/table"
)

type bookReceiptRepository struct {
	cassandraRepository
}

var bookReceiptRepositoryInstance = bookRepository{cassandraRepository{tableDef: table.New(table.Metadata{
	Name:    "book_receipt",
	Columns: ConvertToColumnNames(utils.FieldsOfType((*data.BookReceipt)(nil))),
	PartKey: []string{"store_id"},
	SortKey: []string{"id"},
})}}
