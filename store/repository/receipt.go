package repository

import (
	"store/app/domain"
)

type ReceiptRepository struct {
	postgresRepository[domain.ReceiptData]
}

func (ReceiptRepository) New() *ReceiptRepository {
	return &ReceiptRepository{postgresRepository: postgresRepository[domain.ReceiptData]{eventDispatcher: GetEventDispatcher(), db: GetDb()}}
}

func (r *ReceiptRepository) Create(
	receipt *domain.Receipt,
	tx *Transaction,
) error {
	receiptData := receipt.State()

	if tx == nil {
		tx = Transaction{}.New()
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

	bookRepository := BookRepository{}

	if receiptData.StockAdjustment != nil {
		for _, item := range receiptData.StockAdjustment {
			if err := bookRepository.adjustStock(item, tx); err != nil {
				return err
			}
		}
	}

	return nil
}
