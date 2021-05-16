package domain

import (
	"errors"
	"fmt"
	"store/app/data"
)

type Order struct {
	data.Order
}

func (Order) Get(id data.EntityId, tx data.Transaction) (*Order, error) {
	result, err := OrderRepository.
		Query(&data.Order{}, tx).
		IncludeMany("Items").
		Where("id = ?", id).
		First()

	if err != nil {
		return nil, err
	}

	dataOrder := result.(data.Order)

	return &Order{dataOrder}, nil
}

func (Order) GetReceivingOrders(tx data.Transaction) ([]*Order, error) {
	records, err := OrderRepository.
		Query(&data.Order{}, tx).
		Where("status = ?", data.OrderStatusReceiving).
		IncludeMany("Items").
		OrderBy("created_at").
		Find()

	if err != nil {
		return nil, err
	}

	orders := []*Order{}
	for _, record := range records {
		orders = append(orders, &Order{record.(data.Order)})
	}

	return orders, nil
}

func (order *Order) Accept(tx data.Transaction) error {
	if order.Status != data.OrderStatusQueued && order.Status != data.OrderStatusStockFilled {
		return errors.New(fmt.Sprintf("Order status '%s' is invalid for accepting", order.Status))
	}

	_, err := TransactionFactory.RunInTransaction(func(tx data.Transaction) (interface{}, error) {
		err := OrderRepository.Update(
			order.Id,
			&data.Order{Status: data.OrderStatusAccepted},
			tx,
		)

		if err != nil {
			return nil, err
		}

		// decrease on-hand qty of books associated with the order items
		for _, item := range order.Items {
			book, subErr := Book{}.Get(item.BookId, tx)
			if err = subErr; err != nil {
				return nil, err
			}

			subErr = book.AdjustOnhandQty(-item.Qty, tx)
			if err = subErr; err != nil {
				return nil, err
			}
		}

		return nil, nil
	}, tx)

	return err
}

func (order *Order) PlaceAsBackOrder(tx data.Transaction) error {
	if order.Status != data.OrderStatusQueued {
		return errors.New(
			fmt.Sprintf("Order status '%s' is invalid to be placed as backorder", order.Status),
		)
	}

	_, err := TransactionFactory.RunInTransaction(func(tx data.Transaction) (interface{}, error) {
		err := OrderRepository.Update(order.Id, &data.Order{Status: data.OrderStatusReceiving}, tx)
		if err != nil {
			return nil, err
		}

		// increase preserved qty of books associated with the order items
		for _, item := range order.Items {
			book, subErr := Book{}.Get(item.BookId, tx)
			if err = subErr; err != nil {
				return nil, err
			}

			subErr = book.AdjustPreservedQty(item.Qty, tx)
			if err = subErr; err != nil {
				return nil, err
			}
		}

		return nil, nil
	}, tx)

	return err
}

func (order *Order) TryUpdateToStockFilled(
	stock Stock,
	tx data.Transaction,
) (Stock, error) {
	if order.Status != data.OrderStatusReceiving {
		return stock, errors.New(
			fmt.Sprintf("Order status '%s' is invalid for StockFilled", order.Status),
		)
	}

	if !stock.Enough(order.Items) {
		return stock, errors.New("Not enough stock")
	}

	adjustedStock, err := TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			err := OrderRepository.Update(
				order.Id,
				&data.Order{Status: data.OrderStatusStockFilled},
				tx,
			)

			for _, item := range order.Items {
				book, subErr := Book{}.Get(item.BookId, tx)
				if err = subErr; err != nil {
					return stock, err
				}

				subErr = book.AdjustPreservedQty(-item.Qty, tx)
				if err = subErr; err != nil {
					return stock, err
				}
			}

			if err != nil {
				return stock, err
			}

			return stock.Issue(order.Items), nil
		},
		tx,
	)

	return adjustedStock.(Stock), err
}

func (order *Order) Reject(tx data.Transaction) error {
	if order.Status != data.OrderStatusQueued && order.Status != data.OrderStatusStockFilled {
		return errors.New(fmt.Sprintf("Order status '%s' is invalid for rejecting", order.Status))
	}

	err := OrderRepository.Update(order.Id, &data.Order{Status: data.OrderStatusRejected}, tx)
	if err != nil {
		return err
	}

	return nil
}
