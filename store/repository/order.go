package repository

import (
	"store/app/domain"

	"github.com/thoas/go-funk"
)

type OrderRepository struct {
	postgresRepository[domain.OrderData]
}

func (OrderRepository) New() *OrderRepository {
	return &OrderRepository{postgresRepository: postgresRepository[domain.OrderData]{eventDispatcher: GetEventDispatcher(), db: GetDb()}}
}

func (r *OrderRepository) Get(id string, tx *Transaction) (*domain.Order, error) {
	order, err := r.
		query(tx).
		IncludeMany("Items").ThenInclude("Book").
		Include("Customer").
		Where("id").Eq(id).
		First()

	if err != nil {
		return nil, err
	}

	stock := funk.Map(order.Items, func(item domain.OrderItem) (string, domain.StockItem) {
		return item.BookId, domain.StockItem{
			BookId:      item.BookId,
			OnhandQty:   item.Book.OnhandQty,
			ReservedQty: item.Book.ReservedQty,
		}
	}).(domain.Stock)

	return domain.Order{}.New(order, stock), nil
}

func (r *OrderRepository) Create(order *domain.Order, tx *Transaction) error {
	orderData := order.State()

	if tx == nil {
		tx = Transaction{}.New()
	}

	if err := r.create(orderData.OrderData, tx); err != nil {
		return err
	}

	orderItemRepository := &postgresRepository[domain.OrderItem]{}

	for _, item := range orderData.Items {
		if err := orderItemRepository.create(item, tx); err != nil {
			return err
		}
	}

	if orderData.StockAdjustment != nil {
		bookRepository := &BookRepository{}

		for _, item := range orderData.StockAdjustment {
			if err := bookRepository.adjustStock(item, tx); err != nil {
				return err
			}
		}
	}

	// TODO make sure events are delivered at least once
	go r.eventDispatcher.Dispatch("order", orderData.Id, order.PendingEvents()...)

	return nil
}

func (r *OrderRepository) Update(order *domain.Order, tx *Transaction) error {
	orderData := order.State()

	if tx == nil {
		tx = Transaction{}.New()
	}

	if err := r.update(orderData.OrderData, tx); err != nil {
		return err
	}

	orderItemRepository := &postgresRepository[domain.OrderItem]{}

	if err := orderItemRepository.batchDelete(tx, "order_id = ?", orderData.Id); err != nil {
		return err
	}

	for _, item := range orderData.Items {
		if err := orderItemRepository.create(item, tx); err != nil {
			return err
		}
	}

	if orderData.StockAdjustment != nil {
		bookRepository := &BookRepository{}

		for _, item := range orderData.StockAdjustment {
			if err := bookRepository.adjustStock(item, tx); err != nil {
				return err
			}
		}
	}

	// TODO make sure events are delivered at least once
	go r.eventDispatcher.Dispatch("order", orderData.Id, order.PendingEvents()...)

	return nil
}
