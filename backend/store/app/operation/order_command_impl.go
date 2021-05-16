package operation

import (
	"store/app/data"
	"store/app/domain"
)

type orderCommand struct{}

func (*orderCommand) Accept(id string) error {
	_, err := domain.TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			order, err := domain.Order{}.Get(data.FromStringToEntityId(id), tx)
			if err != nil {
				return nil, err
			}

			if err = order.Accept(tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
		nil,
	)

	return err
}

func (*orderCommand) PlaceAsBackOrder(id string) error {
	_, err := domain.TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			order, err := domain.Order{}.Get(data.FromStringToEntityId(id), tx)
			if err != nil {
				return nil, err
			}

			if err = order.PlaceAsBackOrder(tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
		nil,
	)

	return err
}

func (*orderCommand) Reject(id string) error {
	_, err := domain.TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			order, err := domain.Order{}.Get(data.FromStringToEntityId(id), tx)
			if err != nil {
				return nil, err
			}

			if err = order.Reject(tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
		nil,
	)

	return err
}
