package repository

import (
	"store/app/domain"
	data "store/app/domain/data"
	repo "store/app/repository"

	"github.com/thoas/go-funk"
)

type orderRepository struct {
	postgresRepository
}

func (r *orderRepository) Get(id string, tx repo.Transaction) (*domain.Order, error) {
	record, err := r.
		query(&data.Order{}, tx).
		IncludeMany("Items").
		ThenInclude("Book").
		Where("id").Eq(id).
		First()

	if err != nil {
		return nil, err
	}

	order := record.(data.Order)
	order.Stock = funk.Map(order.Items, func(item data.OrderItem) data.StockItem {
		return data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   item.Book.OnhandQty,
			ReservedQty: item.Book.ReservedQty,
		}
	}).(data.Stock)

	return domain.Order{}.New(order), nil
}

func (r *orderRepository) Update(order *domain.Order, tx repo.Transaction) error {
	dataOrder := order.State()

	if err := r.update(dataOrder.Id, &dataOrder, tx); err != nil {
		return err
	}

	for _, item := range dataOrder.Items {
		if stock, ok := dataOrder.Stock[item.BookId]; ok {
			if err := r.update(
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
