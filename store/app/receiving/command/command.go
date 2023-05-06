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
	receiptRepository := container.Instance().Get(utils.Nameof((*repo.ReceiptRepository)(nil))).(repo.ReceiptRepository)
	transactionFactory := container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			books := []*domain.Book{}

			// create books if not exists
			for _, item := range request.Items {
				book := domain.Book{}.New(item.Book.toDataObject())
				books = append(books, book)

				if err := bookRepository.CreateIfNotExist(book, tx); err != nil {
					return nil, err
				}
			}

			// create book receipt
			receivingBooks := []domain.ReceivingBook{}
			for index, item := range request.Items {
				receivingBooks = append(
					receivingBooks,
					domain.ReceivingBook{BookData: books[index].State(), ReceivingQty: item.Qty},
				)
			}

			newReceipt := domain.Receipt{}.NewFromReceivingBooks(receivingBooks)
			if err := receiptRepository.Create(newReceipt, tx); err != nil {
				return nil, err
			}

			return newReceipt.State().Id, nil
		})

	return err
}
