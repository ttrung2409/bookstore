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
		func(tx *repo.Transaction) (any, error) {
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
			var receivingBooks = make([]struct {
				Id  string
				Qty int
			}, len(request.Items))

			for index, item := range request.Items {
				receivingBooks = append(
					receivingBooks,
					struct {
						Id  string
						Qty int
					}{Id: books[index].State().Id, Qty: item.Qty},
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
