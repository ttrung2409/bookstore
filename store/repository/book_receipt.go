package repository

import (
	"store/app/domain"
	"store/app/domain/data"
	repo "store/app/repository"
)

type bookReceiptRepository struct {
	postgresRepository[data.BookReceipt]
}

func (r *bookReceiptRepository) Get(
	id string,
	tx repo.Transaction,
) (*domain.BookReceipt, error) {
	bookReceipt, err := r.
		query(tx).
		Include("Items").
		ThenInclude("Book").
		Where("id").Eq(id).
		First()

	if err != nil {
		return nil, err
	}

	return domain.BookReceipt{}.New(bookReceipt), nil
}

func (r *bookReceiptRepository) Create(
	receipt *domain.BookReceipt,
	tx repo.Transaction,
) (string, error) {
	dataReceipt := receipt.State()
	dataReceipt.Id = data.NewId()

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

	if err := r.create(dataReceipt, tx); err != nil {
		return data.EmptyId, err
	}

	bookReceiptItemRepository := postgresRepository[data.BookReceiptItem]{}

	for _, item := range dataReceipt.Items {
		if err := bookReceiptItemRepository.create(item, tx); err != nil {
			return data.EmptyId, err
		}
	}

	bookRepository := bookRepository{}

	for _, item := range dataReceipt.OnhandStockAdjustment {
		if err := bookRepository.adjustOnhandQty(item.BookId, item.Qty, tx); err != nil {
			return data.EmptyId, err
		}
	}

	return dataReceipt.Id, nil
}
