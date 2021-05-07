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
	Qty int
}

var bookReceiptRepository = module.Container.Get(utils.Nameof((*data.BookReceiptRepository)(nil))).(data.BookReceiptRepository)

var bookReceiptItemRepository = module.Container.Get(utils.Nameof((*data.BookReceiptItemRepository)(nil))).(data.BookReceiptItemRepository)

var transactionFactory = module.Container.Get(utils.Nameof((*data.TransactionFactory)(nil))).(data.TransactionFactory)

func (BookReceipt) Create(
	books []ReceivingBook,
	transaction *data.Transaction,
) (*BookReceipt, error) {
	receipt := data.BookReceipt{
		Id:      data.NewEntityId(),
		StoreId: data.StoreId(),
	}

	var err error

	if transaction == nil {
		transaction = transactionFactory.New()
	}

	_, err = bookReceiptRepository.Create(receipt, transaction)
	if err != nil {
		return nil, err
	}

	items := []data.BookReceiptItem{}

	for _, book := range books {
		item := data.BookReceiptItem{
			BookId:        book.Id,
			BookReceiptId: receipt.Id,
			Qty:           book.Qty,
		}

		_, err = bookReceiptItemRepository.Create(item, transaction)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if transaction != nil {
		return &BookReceipt{BookReceipt: receipt, Items: items}, nil
	}

	err = (*transaction).Commit()
	if err != nil {
		(*transaction).Rollback()

		return nil, err
	}

	return &BookReceipt{BookReceipt: receipt, Items: items}, nil
}
