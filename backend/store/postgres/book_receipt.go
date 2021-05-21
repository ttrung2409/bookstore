package postgres

import (
	"store/app/data"
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
	tx data.Transaction,
) (*data.BookReceipt, error) {
	record, err := r.
		Query(data.BookReceipt{}, tx).
		Include("Items").
		ThenInclude("Book").
		Where("id == ?", id).
		First()

	if err != nil {
		return nil, err
	}

	receipt := record.(*data.BookReceipt)

	stock := data.Stock{}
	for _, item := range receipt.Items {
		stock[item.BookId] = data.StockItem{
			BookId:      item.BookId,
			OnhandQty:   item.Book.OnhandQty,
			ReservedQty: item.Book.ReservedQty,
		}
	}

	receipt.Stock = stock

	return receipt, nil
}

func (r *bookReceiptRepository) Create(
	receipt *data.BookReceipt,
	tx data.Transaction,
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

	return receiptId, nil
}

// Though updating a receipt only involve updating its stock for the time being,
// this method is still invoked at the receipt level to give the client
// the semantics of updating the receipt as a whole for data consistency
func (r *bookReceiptRepository) Update(receipt *data.BookReceipt, tx data.Transaction) error {
	for _, item := range receipt.Items {
		if stock, ok := receipt.Stock[item.BookId]; ok {
			if err := bookRepositoryInstance.update(
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
