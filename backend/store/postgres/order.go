package postgres

import (
	data "store/app/domain/data"
	repo "store/app/repository"

	"github.com/thoas/go-funk"
)

type orderRepository struct {
	postgresRepository
}

func (r *orderRepository) Get(id data.EntityId, tx repo.Transaction) (*data.Order, error) {
	record, err := r.
		Query(&data.Order{}, tx).
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id = ?", id).
		First()

	if err != nil {
		return nil, err
	}

	order := record.(*data.Order)
	order.Stock = funk.Map(order.Items, func(item data.OrderItem) data.StockItem {
		return data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   item.Book.OnhandQty,
			ReservedQty: item.Book.ReservedQty,
		}
	}).(data.Stock)

	return order, nil
}

func (r *orderRepository) GetReceivingOrders(tx repo.Transaction) ([]*data.Order, error) {
	records, err := r.
		Query(&data.Order{}, tx).
		Where("status = ?", data.OrderStatusReceiving).
		IncludeMany("Items").
		ThenInclude("Book").
		OrderBy("created_at").
		Find()

	if err != nil {
		return nil, err
	}

	orders := []*data.Order{}
	for _, record := range records {
		order := record.(*data.Order)
		order.Stock = funk.Map(order.Items, func(item data.OrderItem) data.StockItem {
			return data.StockItem{
				BookId:      item.BookId,
				OnhandQty:   item.Book.OnhandQty,
				ReservedQty: item.Book.ReservedQty,
			}
		}).(data.Stock)

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) Update(order *data.Order, tx repo.Transaction) error {
	if err := r.update(order.Id, order, tx); err != nil {
		return err
	}

	for _, item := range order.Items {
		if stock, ok := order.Stock[item.BookId]; ok {
			if err := bookRepositoryInstance.update(
				stock.BookId,
				&data.Book{OnhandQty: stock.OnhandQty, ReservedQty: stock.ReservedQty},
				tx,
			); err != nil {
				return err
			}
		}
	}

	return nil
}

var orderRepositoryInstance = orderRepository{postgresRepository{newEntity: func() data.Entity {
	return &data.Order{}
}}}
