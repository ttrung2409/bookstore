package command

import (
	"store/app/domain"
	"store/app/messaging"
	"store/app/messaging/events"
	repo "store/app/repository"
	"store/container"
	"store/utils"

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
	transactionFactory := container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)
	orderRepository := container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)
	bookRepository := container.Instance().Get(utils.Nameof((*repo.BookRepository)(nil))).(repo.BookRepository)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			dataOrder := order.toDataObject()
			dataOrder.Stock = bookRepository.GetStock(funk.Map(order.Items, func(item OrderItem) string {
				return item.BookId
			}).([]string))

			order := domain.Order{}.New(dataOrder)

			if err := order.Accept(); err != nil {
				return nil, err
			}

			if _, err := orderRepository.Create(order, tx); err != nil {
				return nil, err
			}

			return nil, nil
		},
	)

	if err != nil {
		return err
	}

	eventDispatcher := container.Instance().Get(utils.Nameof((*messaging.EventDispatcher)(nil))).(messaging.EventDispatcher)
	if err := eventDispatcher.Dispatch(&events.OrderAccepted{OrderId: order.Id}); err != nil {
		return err
	}

	return nil
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
