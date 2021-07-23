package command

import (
	"store/app/data"
	"store/app/domain"
)

type acceptOrderCommand struct{}

func (*acceptOrderCommand) Execute(id string) error {
	_, err := TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			dataOrder, err := OrderRepository.Get(data.FromStringToEntityId(id), tx)
			if err != nil {
				return nil, err
			}

			order := domain.Order{}.New(dataOrder)

			if err = order.Accept(); err != nil {
				return nil, err
			}

			if err = OrderRepository.Update(order.State(), tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
		nil,
	)

	return err
}
