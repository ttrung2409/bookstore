package postgres

import (
	"store/app/data"
	"store/utils"

	"github.com/sarulabs/di"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Install(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*data.BookRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return bookRepositoryInstance, nil
			},
		},
		{
			Name: utils.Nameof((*data.BookReceiptRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return bookReceiptRepositoryInstance, nil
			},
		},
		{
			Name: utils.Nameof((*data.BookReceiptItemRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return bookReceiptItemRepositoryInstance, nil
			},
		},
		{
			Name: utils.Nameof((*data.OrderRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return orderRepositoryInstance, nil
			},
		},
		{
			Name: utils.Nameof((*data.TransactionFactory)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return transactionFactory{}, nil
			},
		},
	}...)
}

func Db() *gorm.DB {
	if db == nil {
		db = connect()
	}

	return db
}

func connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=bookstore port=5432"
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			}},
	)

	if err != nil {
		panic(err)
	}

	return db
}
