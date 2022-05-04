package command

import (
	"store/app/domain"
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

type Command interface {
	Receive(request ReceiveBooksRequest) error
}

func New() Command {
	return &command{}
}

type command struct{}

func (*command) Receive(request ReceiveBooksRequest) error {
	bookRepository := container.Instance().Get(utils.Nameof((*repo.BookRepository)(nil))).(repo.BookRepository)
	bookReceiptRepository := container.Instance().Get(utils.Nameof((*repo.BookReceiptRepository)(nil))).(repo.BookReceiptRepository)
	transactionFactory := container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			bookIds := make([]string, len(request.Items))

			// create books if not exists
			for index, item := range request.Items {
				bookId, err := bookRepository.CreateIfNotExist(
					domain.Book{}.New(item.Book.toDataObject()),
					tx,
				)

				if err != nil {
					return nil, err
				}

				bookIds[index] = bookId
			}

			// create book receipt
			receivingBooks := []domain.ReceivingBook{}
			for index, item := range request.Items {
				dataBook := item.toDataObject()
				dataBook.Id = bookIds[index]

				receivingBooks = append(
					receivingBooks,
					domain.ReceivingBook{Book: dataBook, ReceivingQty: item.Qty},
				)
			}

			newReceipt := domain.BookReceipt{}.NewFromReceivingBooks(receivingBooks)
			receiptId, err := bookReceiptRepository.Create(newReceipt, tx)
			if err != nil {
				return nil, err
			}

			return receiptId, err
		})

	return err
}
