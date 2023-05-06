package repository

import (
	"store/app/domain/data"
	repo "store/app/repository"
	"store/utils"

	"github.com/sarulabs/di"
)

func RegisterDependencies(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*repo.BookRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return &bookRepository{postgresRepository[data.Book]{eventDispatcher: GetEventDispatcher(), db: Db()}}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.ReceiptRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return &receiptRepository{postgresRepository[data.Receipt]{eventDispatcher: GetEventDispatcher(), db: Db()}}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.OrderRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return &orderRepository{postgresRepository[data.Order]{eventDispatcher: GetEventDispatcher(), db: Db()}}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.TransactionFactory)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return transactionFactory{}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.QueryFactory[data.Book])(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return queryFactory[data.Book]{}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.QueryFactory[data.Order])(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return queryFactory[data.Order]{}, nil
			},
		},
	}...)
}
