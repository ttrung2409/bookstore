package domain

import (
	module "store"
	data "store/app/data"
	"store/utils"
)

type BookReceipt struct {
	data.BookReceipt
}

type ReceivingBook struct {
	Book data.Book
	Qty  int
}

var bookReceiptRepository = module.Container.Get(utils.Nameof((*data.BookReceiptRepository)(nil))).(data.BookReceiptRepository)

var bookReceiptItemRepository = module.Container.Get(utils.Nameof((*data.BookReceiptItemRepository)(nil))).(data.BookReceiptItemRepository)

var transactionFactory = module.Container.Get(utils.Nameof((*data.TransactionFactory)(nil))).(data.TransactionFactory)

func (BookReceipt) Create(
	items []ReceivingBook,
	transaction *data.Transaction,
) (*BookReceipt, error) {
	receipt := data.BookReceipt{
		Id: data.NewEntityId(),
	}

	var err error

	if transaction == nil {
		transaction = transactionFactory.New()
	}

	_, err = bookReceiptRepository.Create(receipt, transaction)

	for _, item := range items {
		_, err = bookReceiptItemRepository.Create(data.BookReceiptItem{
			BookId:        item.Book.Id.Value(),
			BookReceiptId: receipt.Id,
			Qty:           item.Qty,
		}, transaction)
	}

	if transaction != nil {
		return &BookReceipt{receipt}, err
	}

	err = (*transaction).Commit()
	if err != nil {
		(*transaction).Rollback()

		return nil, err
	}

	return &BookReceipt{receipt}, nil
}
