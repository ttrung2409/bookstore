package repository

import (
	"store/app/domain"
	repo "store/app/repository"

	"github.com/thoas/go-funk"
)

type orderRepository struct {
	postgresRepository[domain.OrderData]
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

	stock := funk.Map(order.Items, func(item domain.OrderItemData) (string, domain.StockItemData) {
		return item.BookId, domain.StockItemData{
			BookId:      item.BookId,
			OnhandQty:   item.Book.OnhandQty,
			ReservedQty: item.Book.ReservedQty,
		}
	}).(domain.StockData)

	return domain.Order{}.New(order, stock), nil
}

func (r *orderRepository) Create(order *domain.Order, tx repo.Transaction) error {
	orderData := order.State()

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

	if err := r.create(orderData.OrderData, tx); err != nil {
		return err
	}

	orderItemRepository := &postgresRepository[domain.OrderItemData]{}

	for _, item := range orderData.Items {
		if err := orderItemRepository.create(item, tx); err != nil {
			return err
		}
	}

	bookRepository := &bookRepository{}

	for _, item := range orderData.Items {
		if stock, ok := orderData.Stock[item.BookId]; ok {
			if err := bookRepository.update(
				domain.BookData{Id: stock.BookId, OnhandQty: stock.OnhandQty, ReservedQty: stock.ReservedQty},
				tx,
			); err != nil {
				return err
			}
		}
	}

	// TODO make sure events are delivered at least once
	go r.eventDispatcher.Dispatch("order", orderData.Id, order.PendingEvents...)

	return nil
}

func (r *orderRepository) Update(order *domain.Order, tx repo.Transaction) error {
	orderData := order.State()

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

	if err := r.update(orderData.OrderData, tx); err != nil {
		return err
	}

	orderItemRepository := &postgresRepository[domain.OrderItemData]{}

	if err := orderItemRepository.batchDelete(tx, "order_id = ?", orderData.Id); err != nil {
		return err
	}

	for _, item := range orderData.Items {
		if err := orderItemRepository.create(item, tx); err != nil {
			return err
		}
	}

	bookRepository := &bookRepository{}

	for _, item := range orderData.Items {
		if stock, ok := orderData.Stock[item.BookId]; ok {
			if err := bookRepository.update(
				domain.BookData{Id: stock.BookId, OnhandQty: stock.OnhandQty, ReservedQty: stock.ReservedQty},
				tx,
			); err != nil {
				return err
			}
		}
	}

	// TODO make sure events are delivered at least once
	go r.eventDispatcher.Dispatch("order", orderData.Id, order.PendingEvents...)

	return nil
}
