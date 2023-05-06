package repository

import (
	"store/app/domain"
	repo "store/app/repository"
	"store/utils"

	"github.com/sarulabs/di"
)

func RegisterDependencies(builder *di.Builder) {
	builder.Add([]di.Def{
		{
			Name: utils.Nameof((*repo.BookRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return &bookRepository{postgresRepository[domain.BookData]{eventDispatcher: GetEventDispatcher(), db: Db()}}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.ReceiptRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return &receiptRepository{postgresRepository[domain.ReceiptData]{eventDispatcher: GetEventDispatcher(), db: Db()}}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.OrderRepository)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return &orderRepository{postgresRepository[domain.OrderData]{eventDispatcher: GetEventDispatcher(), db: Db()}}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.TransactionFactory)(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return transactionFactory{}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.QueryFactory[domain.BookData])(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return queryFactory[domain.BookData]{}, nil
			},
		},
		{
			Name: utils.Nameof((*repo.QueryFactory[domain.OrderData])(nil)),
			Build: func(ctn di.Container) (interface{}, error) {
				return queryFactory[domain.OrderData]{}, nil
			},
		},
	}...)
}
