package domain

import (
	data "store/app/data"
)

type BookReceipt struct {
	data.BookReceipt
}

type ReceivingBook struct {
	data.Book
	ReceivingQty int
}

func (BookReceipt) Create(
	books []ReceivingBook,
	tx data.Transaction,
) (*BookReceipt, error) {
	receipt, err := TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			receipt := data.BookReceipt{
				Id: data.NewEntityId(),
			}

			_, err := BookReceiptRepository.Create(&receipt, tx)
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

			receipt.Items = items

			return &BookReceipt{receipt}, nil
		},
		tx,
	)

	return receipt.(*BookReceipt), err
}

func (receipt *BookReceipt) IncreaseStock(tx data.Transaction) (Stock, error) {
	stock, err := TransactionFactory.RunInTransaction(func(tx data.Transaction) (interface{}, error) {
		stock := Stock{}
		for _, item := range receipt.Items {
			book, err := Book{}.Get(item.BookId, tx)
			if err != nil {
				return nil, err
			}

			err = book.AdjustOnhandQty(item.Qty, tx)
			if err != nil {
				return nil, err
			}

			stock[item.BookId] = data.StockItem{
				BookId:       book.Id,
				OnhandQty:    book.OnhandQty + item.Qty,
				PreservedQty: book.PreservedQty,
			}
		}

		return stock, nil
	}, tx)

	return stock.(Stock), err
}
