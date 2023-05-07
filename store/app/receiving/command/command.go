package command

import (
	"store/app/domain"
	repo "store/repository"
)

type Command interface {
	Receive(request ReceiveBooksRequest) error
}

func New() Command {
	return &command{}
}

type command struct{}

func (*command) Receive(request ReceiveBooksRequest) error {
	bookRepository := repo.BookRepository{}.New()
	receiptRepository := repo.ReceiptRepository{}.New()

	_, err := repo.Transaction{}.RunInTransaction(
		func(tx *repo.Transaction) (interface{}, error) {
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
