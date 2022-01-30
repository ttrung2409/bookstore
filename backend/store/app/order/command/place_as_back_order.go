package command

import (
	"store/app/domain/data"
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

func (*command) PlaceAsBackOrder(orderId string) error {
	var transactionFactory = container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
	var orderRepository = container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			order, err := orderRepository.Get(data.FromStringToEntityId(orderId), tx)
			if err != nil {
				return nil, err
			}

			if err = order.PlaceAsBackOrder(); err != nil {
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
