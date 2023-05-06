package repository

import (
	"store/app/domain"
	"store/app/domain/data"
	repo "store/app/repository"
)

type receiptRepository struct {
	postgresRepository[data.Receipt]
}

func (r *receiptRepository) Get(
	id string,
	tx repo.Transaction,
) (*domain.Receipt, error) {
	receipt, err := r.
		query(tx).
		Include("Items").
		ThenInclude("Book").
		Where("id").Eq(id).
		First()

	if err != nil {
		return nil, err
	}

	return domain.Receipt{}.New(receipt), nil
}

func (r *receiptRepository) Create(
	receipt *domain.Receipt,
	tx repo.Transaction,
) error {
	dataReceipt := receipt.State()

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

	if err := r.create(dataReceipt, tx); err != nil {
		return err
	}

	receiptItemRepository := postgresRepository[data.ReceiptItem]{}

	for _, item := range dataReceipt.Items {
		if err := receiptItemRepository.create(item, tx); err != nil {
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
