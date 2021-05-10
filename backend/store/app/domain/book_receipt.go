package domain

import (
	module "store"
	data "store/app/data"
	"store/utils"
)

type BookReceipt struct {
	data.BookReceipt
	Items []data.BookReceiptItem
}

type ReceivingBook struct {
	data.Book
	ReceivingQty int
}

var bookReceiptRepository = module.Container().Get(utils.Nameof((*data.BookReceiptRepository)(nil))).(data.BookReceiptRepository)

var bookReceiptItemRepository = module.Container().Get(utils.Nameof((*data.BookReceiptItemRepository)(nil))).(data.BookReceiptItemRepository)

var transactionFactory = module.Container().Get(utils.Nameof((*data.TransactionFactory)(nil))).(data.TransactionFactory)

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
		tx = transactionFactory.New()
	}

	_, err = bookReceiptRepository.Create(receipt, tx)
	if err != nil {
		return nil, err
	}

	items := []data.BookReceiptItem{}

	for _, book := range books {
		item := data.BookReceiptItem{
			BookReceiptId: receipt.Id,
			BookId:        book.Id,
			Qty:           book.ReceivingQty,
		}

		_, err = bookReceiptItemRepository.Create(item, tx)
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
