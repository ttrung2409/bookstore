package repository

import (
	"store/app/domain"
	"store/app/domain/data"
	repo "store/app/repository"
)

type bookReceiptRepository struct {
	postgresRepository
}

type bookReceiptItemRepository struct {
	postgresRepository
}

func (r *bookReceiptRepository) Get(
	id string,
	tx repo.Transaction,
) (*domain.BookReceipt, error) {
	record, err := r.
		query(&data.BookReceipt{}, tx).
		Include("Items").
		ThenInclude("Book").
		Where("id").Eq(id).
		First()

	if err != nil {
		return nil, err
	}

	return domain.BookReceipt{}.New(record.(*data.BookReceipt)), nil
}

func (r *bookReceiptRepository) Create(
	receipt *domain.BookReceipt,
	tx repo.Transaction,
) (string, error) {
	dataReceipt := receipt.State()

	receiptId, err := r.create(dataReceipt, tx)
	if err != nil {
		return data.EmptyEntityId, nil
	}

	bookReceiptItemRepositoryInstance := bookReceiptItemRepository{postgresRepository{}}

	for _, item := range dataReceipt.Items {
		if _, err = bookReceiptItemRepositoryInstance.create(item, tx); err != nil {
			return data.EmptyEntityId, err
		}
	}

	bookRepositoryInstance := bookRepository{postgresRepository{}}

	for _, item := range dataReceipt.OnhandStockAdjustment {
		bookRepositoryInstance.adjustOnhandQty(item.BookId, item.Qty, tx)
	}

	return receiptId, nil
}
