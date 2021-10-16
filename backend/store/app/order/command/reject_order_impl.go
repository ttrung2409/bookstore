package command

import (
	"store/app/domain"
	"store/app/domain/data"
	repo "store/app/repository"
)

type rejectOrderCommand struct{}

func (*rejectOrderCommand) Execute(id string) error {
	_, err := TransactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			dataOrder, err := OrderRepository.Get(data.FromStringToEntityId(id), tx)
			if err != nil {
				return nil, err
			}

			order := domain.Order{}.New(dataOrder)

			if err = order.Reject(); err != nil {
				return nil, err
			}

			if err = OrderRepository.Update(order.State(), tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
	)

	return err
}
