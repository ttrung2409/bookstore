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
) error {
	dataReceipt := receipt.State()

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

	if err := r.create(dataReceipt, tx); err != nil {
		return err
	}

	bookReceiptItemRepository := postgresRepository[data.BookReceiptItem]{}

	for _, item := range dataReceipt.Items {
		if err := bookReceiptItemRepository.create(item, tx); err != nil {
			return err
		}
	}

	bookRepository := bookRepository{}

	for _, item := range dataReceipt.OnhandStockAdjustment {
		if err := bookRepository.adjustOnhandQty(item.BookId, item.Qty, tx); err != nil {
			return err
		}
	}

	return nil
}
