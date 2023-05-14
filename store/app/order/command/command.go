package command

import (
	"store/app/domain"
	"store/kafka"
	repo "store/repository"

	"github.com/thoas/go-funk"
)

type Command interface {
	AcceptOrder(order Order) error
	DeliverOrder(orderId string) error
	CancelOrder(orderId string) error
}

func New() Command {
	return &command{}
}

type command struct{}

func (*command) AcceptOrder(order Order) error {
	bookRepository := repo.BookRepository{}.New()
	orderRepository := repo.OrderRepository{}.New()

	_, err := repo.Transaction{}.RunInTransaction(
		func(tx *repo.Transaction) (any, error) {
			stock, err := bookRepository.GetStock(funk.Map(order.Items, func(item OrderItem) string {
				return item.BookId
			}).([]string), tx)

			if err != nil {
				return nil, err
			}

			order := domain.Order{}.New(order.toDataObject(), stock)

			if err := order.Accept(); err != nil {
				return nil, err
			}

			if err := orderRepository.Create(order, tx); err != nil {
				return nil, err
			}

			// TODO make sure events are delivered at least once
			go kafka.GetEventDispatcher().Dispatch("order", order.State().Id, order.PendingEvents()...)

			return nil, nil
		},
	)

	return err
}

func (*command) CancelOrder(orderId string) error {
	orderRepository := repo.OrderRepository{}.New()

	_, err := repo.Transaction{}.RunInTransaction(
		func(tx *repo.Transaction) (any, error) {
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
	orderRepository := repo.OrderRepository{}.New()

	_, err := repo.Transaction{}.RunInTransaction(
		func(tx *repo.Transaction) (any, error) {
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

			// TODO make sure events are delivered at least once
			go kafka.GetEventDispatcher().Dispatch("order", order.State().Id, order.PendingEvents()...)

			return nil, nil
		},
	)

	return err
}
