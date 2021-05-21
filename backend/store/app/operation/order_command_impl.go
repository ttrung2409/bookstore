package operation

import (
	"store/app/data"
	"store/app/domain"
)

type orderCommand struct{}

func (*orderCommand) Accept(id string) error {
	_, err := TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			dataOrder, err := OrderRepository.Get(data.FromStringToEntityId(id), tx)
			if err != nil {
				return nil, err
			}

			order := domain.Order{}.New(dataOrder)

			if err = order.Accept(tx); err != nil {
				return nil, err
			}

			if err = OrderRepository.Update(&order.Order, tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
		nil,
	)

	return err
}

func (*orderCommand) PlaceAsBackOrder(id string) error {
	_, err := TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			dataOrder, err := OrderRepository.Get(data.FromStringToEntityId(id), tx)
			if err != nil {
				return nil, err
			}

			order := domain.Order{}.New(dataOrder)

			if err = order.PlaceAsBackOrder(tx); err != nil {
				return nil, err
			}

			if err = OrderRepository.Update(&order.Order, tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
		nil,
	)

	return err
}

func (*orderCommand) Reject(id string) error {
	_, err := TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			dataOrder, err := OrderRepository.Get(data.FromStringToEntityId(id), tx)
			if err != nil {
				return nil, err
			}

			order := domain.Order{}.New(dataOrder)

			if err = order.Reject(tx); err != nil {
				return nil, err
			}

			if err = OrderRepository.Update(&order.Order, tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
		nil,
	)

	return err
}
