package command

import (
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

type Command interface {
	AcceptOrder(orderId string) error
	DeliverOrder(orderId string) error
	CancelOrder(orderId string) error
}

func New() Command {
	return &command{}
}

type command struct{}

func (*command) AcceptOrder(orderId string) error {
	transactionFactory := container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
	orderRepository := container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			order, err := orderRepository.Get(orderId, tx)
			if err != nil {
				return nil, err
			}

			if err = order.Accept(); err != nil {
				return nil, err
			}

			if err = orderRepository.Update(order, tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
	)

	return err
}

func (*command) CancelOrder(orderId string) error {
	transactionFactory := container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
	orderRepository := container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			order, err := orderRepository.Get(orderId, tx)
			if err != nil {
				return nil, err
			}

			if err = order.Cancel(); err != nil {
				return nil, err
			}

			if err = orderRepository.Update(order, tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
	)

	return err
}

func (*command) DeliverOrder(orderId string) error {
	transactionFactory := container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
	orderRepository := container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			order, err := orderRepository.Get(orderId, tx)
			if err != nil {
				return nil, err
			}

			if err = order.Deliver(); err != nil {
				return nil, err
			}

			if err = orderRepository.Update(order, tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
	)

	return err
}
