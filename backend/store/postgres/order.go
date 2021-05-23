package postgres

import (
	data "store/app/data"
)

type orderRepository struct {
	postgresRepository
}

func (r *orderRepository) Get(id data.EntityId, tx data.Transaction) (*data.Order, error) {
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

	stock := data.Stock{}
	for _, item := range order.Items {
		stock[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   item.Book.OnhandQty,
			ReservedQty: item.Book.ReservedQty,
		}
	}

	order.Stock = stock

	return order, nil
}

func (r *orderRepository) GetReceivingOrders(tx data.Transaction) ([]*data.Order, error) {
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

		stock := data.Stock{}
		for _, item := range order.Items {
			stock[item.BookId] = data.StockItem{
				BookId:      item.BookId,
				OnhandQty:   item.Book.OnhandQty,
				ReservedQty: item.Book.ReservedQty,
			}
		}

		order.Stock = stock
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) Update(order *data.Order, tx data.Transaction) error {
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
