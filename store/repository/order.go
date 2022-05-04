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
		IncludeMany("Items").ThenInclude("Book").
		Include("Customer").
		Where("id").Eq(id).
		First()

	if err != nil {
		return nil, err
	}

	order := record.(data.Order)
	order.Stock = funk.Map(order.Items, func(item data.OrderItem) (string, data.StockItem) {
		return item.BookId, data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   item.Book.OnhandQty,
			ReservedQty: item.Book.ReservedQty,
		}
	}).(data.Stock)

	return domain.Order{}.New(order), nil
}

func (r *orderRepository) Create(order *domain.Order, tx repo.Transaction) (string, error) {
	dataOrder := order.State()

	if dataOrder.Id == "" {
		dataOrder.Id = data.NewEntityId()
	}

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

	if err := r.create(&dataOrder, tx); err != nil {
		return data.EmptyEntityId, err
	}

	for _, item := range dataOrder.Items {
		if err := r.create(&item, tx); err != nil {
			return data.EmptyEntityId, err
		}
	}

	for _, item := range dataOrder.Items {
		if stock, ok := dataOrder.Stock[item.BookId]; ok {
			if err := r.update(
				stock.BookId,
				&data.Book{OnhandQty: stock.OnhandQty, ReservedQty: stock.ReservedQty},
				tx,
			); err != nil {
				return data.EmptyEntityId, err
			}
		}
	}

	return dataOrder.Id, nil
}

func (r *orderRepository) Update(order *domain.Order, tx repo.Transaction) error {
	dataOrder := order.State()

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

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
