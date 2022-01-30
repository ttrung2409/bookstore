package repository

import (
	repo "store/app/repository"
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
			Name: utils.Nameof((*repo.BookRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return bookRepository{postgresRepository{}}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.BookReceiptRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return bookReceiptRepository{postgresRepository{}}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.OrderRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return orderRepository{postgresRepository{}}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.TransactionFactory)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return transactionFactory{}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.QueryFactory)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return queryFactory{}, nil
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