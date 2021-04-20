package domain

import (
	module "store"
	data "store/data"
	"store/utils"
)

type BookReceipt struct {
	data.BookReceipt
}

var bookReceiptRepository = module.Container.Get(utils.Nameof((*data.BookReceiptRepository)(nil))).(data.BookReceiptRepository)

var bookReceiptItemRepository = module.Container.Get(utils.Nameof((*data.BookReceiptItemRepository)(nil))).(data.BookReceiptItemRepository)

var transactionFactory = module.Container.Get(utils.Nameof((*data.TransactionFactory)(nil))).(data.TransactionFactory)

func CreateBookReceipt(
	books []data.Book,
	qty map[string]int,
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

	for _, book := range books {
		_, err = bookReceiptItemRepository.Create(data.BookReceiptItem{
			BookId:        book.Id.Value(),
			BookReceiptId: receipt.Id,
			Qty:           qty[book.Id.Value().ToString()],
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
