package postgres

import (
	"store/app/domain/data"
	repo "store/app/repository"
)

type bookReceiptRepository struct {
	postgresRepository
}

type bookReceiptItemRepository struct {
	postgresRepository
}

var bookReceiptRepositoryInstance = bookReceiptRepository{postgresRepository{newEntity: func() data.Entity {
	return &data.BookReceipt{}
}}}

var bookReceiptItemRepositoryInstance = bookReceiptItemRepository{postgresRepository{newEntity: func() data.Entity {
	return &data.BookReceiptItem{}
}}}

func (r *bookReceiptRepository) Get(
	id data.EntityId,
	tx repo.Transaction,
) (*data.BookReceipt, error) {
	record, err := r.
		Query(&data.BookReceipt{}, tx).
		Include("Items").
		ThenInclude("Book").
		Where("id == ?", id).
		First()

	if err != nil {
		return nil, err
	}

	return record.(*data.BookReceipt), nil
}

func (r *bookReceiptRepository) Create(
	receipt *data.BookReceipt,
	tx repo.Transaction,
) (data.EntityId, error) {
	receiptId, err := r.create(receipt, tx)
	if err != nil {
		return data.EmptyEntityId, nil
	}

	for _, item := range receipt.Items {
		if _, err = bookReceiptItemRepositoryInstance.create(&item, tx); err != nil {
			return data.EmptyEntityId, err
		}
	}

	for _, item := range receipt.OnhandStockAdjustment {
		bookRepositoryInstance.AdjustOnhandQty(item.BookId, item.Qty, tx)
	}

	return receiptId, nil
}
