package repository

import (
	"store/app/domain"
	repo "store/app/repository"
)

type receiptRepository struct {
	postgresRepository[domain.ReceiptData]
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
	receiptData := receipt.State()

	if tx == nil {
		tx = (&transactionFactory{}).New()
	}

	if err := r.create(receiptData.ReceiptData, tx); err != nil {
		return err
	}

	receiptItemRepository := postgresRepository[domain.ReceiptItemData]{}

	for _, item := range receiptData.Items {
		if err := receiptItemRepository.create(item, tx); err != nil {
			return err
		}
	}

	bookRepository := bookRepository{}

	for _, item := range receiptData.OnhandStockAdjustment {
		if err := bookRepository.adjustOnhandQty(item.BookId, item.Qty, tx); err != nil {
			return err
		}
	}

	return nil
}
