package repository

import (
	"store/app/domain"
	repo "store/app/repository"
)

type receiptRepository struct {
	postgresRepository[domain.ReceiptData]
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

	receiptItemRepository := postgresRepository[domain.ReceiptItem]{}

	for _, item := range receiptData.Items {
		if err := receiptItemRepository.create(item, tx); err != nil {
			return err
		}
	}

	bookRepository := bookRepository{}

	if receiptData.StockAdjustment != nil {
		for _, item := range receiptData.StockAdjustment {
			if err := bookRepository.adjustStock(item, tx); err != nil {
				return err
			}
		}
	}

	return nil
}
