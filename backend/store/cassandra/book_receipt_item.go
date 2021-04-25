package cassandra

import (
	data "store/app/data"
	"store/utils"

	"github.com/scylladb/gocqlx/v2/table"
)

type bookReceiptItemRepository struct {
	cassandraRepository
}

var bookReceiptItemRepositoryInstance = &bookRepository{cassandraRepository{tableDef: table.New(table.Metadata{
	Name:    "book_receipt_item",
	Columns: ConvertToColumnNames(utils.FieldsOfType((*data.BookReceiptItem)(nil))),
	PartKey: []string{"book_receipt_id"},
})}}
