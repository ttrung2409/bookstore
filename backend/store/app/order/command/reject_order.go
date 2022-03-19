package command

import (
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

func (*command) RejectOrder(orderId string) error {
	transactionFactory := container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
	orderRepository := container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			order, err := orderRepository.Get(orderId, tx)
			if err != nil {
				return nil, err
			}

			if err = order.Reject(); err != nil {
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
