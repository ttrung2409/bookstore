package repository

import (
	"store/app/domain"
	data "store/app/domain/data"
	repo "store/app/repository"

	"github.com/thoas/go-funk"
)

type orderRepository struct {
	postgresRepository[data.Order]
}

func (r *orderRepository) Get(id string, tx repo.Transaction) (*domain.Order, error) {
	order, err := r.
		query(tx).
		IncludeMany("Items").ThenInclude("Book").
		Include("Customer").
		Where("id").Eq(id).
		First()

	if err != nil {
		return nil, err
	}

	order.Stock = funk.Map(order.Items, func(item data.OrderItem) (string, data.StockItem) {
		return item.BookId, data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   item.Book.OnhandQty,
			ReservedQty: item.Book.ReservedQty,
		}
	}).(data.Stock)

	return domain.Order{}.New(order), nil
}

func (r *orderRepository) Create(order *domain.Order, tx repo.Transaction) error {
	dataOrder := order.State()

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

	if err := r.create(dataOrder, tx); err != nil {
		return err
	}

	orderItemRepository := &postgresRepository[data.OrderItem]{}

	for _, item := range dataOrder.Items {
		if err := orderItemRepository.create(item, tx); err != nil {
			return err
		}
	}

	bookRepository := &bookRepository{}

	for _, item := range dataOrder.Items {
		if stock, ok := dataOrder.Stock[item.BookId]; ok {
			if err := bookRepository.update(
				data.Book{Id: stock.BookId, OnhandQty: stock.OnhandQty, ReservedQty: stock.ReservedQty},
				tx,
			); err != nil {
				return err
			}
		}
	}

	// TODO make sure events are delivered at least once
	go r.eventDispatcher.Dispatch("order", dataOrder.Id, order.PendingEvents...)

	return nil
}

func (r *orderRepository) Update(order *domain.Order, tx repo.Transaction) error {
	dataOrder := order.State()

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

	if err := r.update(dataOrder, tx); err != nil {
		return err
	}

	orderItemRepository := &postgresRepository[data.OrderItem]{}

	if err := orderItemRepository.batchDelete(tx, "order_id = ?", dataOrder.Id); err != nil {
		return err
	}

	for _, item := range dataOrder.Items {
		if err := orderItemRepository.create(item, tx); err != nil {
			return err
		}
	}

	bookRepository := &bookRepository{}

	for _, item := range dataOrder.Items {
		if stock, ok := dataOrder.Stock[item.BookId]; ok {
			if err := bookRepository.update(
				data.Book{Id: stock.BookId, OnhandQty: stock.OnhandQty, ReservedQty: stock.ReservedQty},
				tx,
			); err != nil {
				return err
			}
		}
	}

	// TODO make sure events are delivered at least once
	go r.eventDispatcher.Dispatch("order", dataOrder.Id, order.PendingEvents...)

	return nil
}
