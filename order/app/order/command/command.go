package command

import (
	"ecommerce/app/domain"
	repo "ecommerce/app/repository"
	"ecommerce/app/repository/queries"
	"ecommerce/container"
	"ecommerce/utils"
)

type CreateOrderRequest struct {
	Customer Customer
	BookIds  []string
}

type Command interface {
	CreateOrder(request CreateOrderRequest) error
	CancelOrder(orderId string) error
}

type command struct{}

func (*command) CreateOrder(request CreateOrderRequest) error {
	transactionFactory := container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
	orderRepository := container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)
	booksQuery := container.Instance().Get(utils.Nameof((*queries.BooksQuery)(nil))).(queries.BooksQuery)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (any, error) {

			order := domain.Order.New(request.Customer.toDataObject(), request.Books)
			return nil, nil
		},
	)

	return err
}
