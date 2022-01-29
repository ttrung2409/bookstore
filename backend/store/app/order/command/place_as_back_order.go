package command

import (
	"store/app/domain/data"
	repo "store/app/repository"
)

func (*command) PlaceAsBackOrder(orderId string) error {
	_, err := TransactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			order, err := OrderRepository.Get(data.FromStringToEntityId(orderId), tx)
			if err != nil {
				return nil, err
			}

			if err = order.PlaceAsBackOrder(); err != nil {
				return nil, err
			}

			if err = OrderRepository.Update(order, tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
	)

	return err
}
