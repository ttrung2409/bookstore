package domain

import (
	data "store/app/data"
)

type BookReceipt struct {
	data.BookReceipt
	Items []data.BookReceiptItem
}

type ReceivingBook struct {
	data.Book
	ReceivingQty int
}

func (BookReceipt) Create(
	books []ReceivingBook,
	ambientTx data.Transaction,
) (*BookReceipt, error) {
	receipt := data.BookReceipt{
		Id: data.NewEntityId(),
	}

	var err error

	tx := ambientTx
	if tx == nil {
		tx = TransactionFactory.New()
	}

	_, err = BookReceiptRepository.Create(&receipt, tx)
	if err != nil {
		return nil, err
	}

	items := []data.BookReceiptItem{}

	for _, book := range books {
		item := data.BookReceiptItem{
			Id:            data.NewEntityId(),
			BookReceiptId: receipt.Id,
			BookId:        book.Id,
			Qty:           book.ReceivingQty,
		}

		_, err = BookReceiptItemRepository.Create(&item, tx)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if ambientTx == nil {
		err = tx.Commit()
		if err != nil {
			tx.Rollback()

			return nil, err
		}
	}

	return &BookReceipt{BookReceipt: receipt, Items: items}, nil
}
