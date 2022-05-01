package command

import (
	"store/app/domain"
	repo "store/app/repository"
	"store/container"
	"store/utils"

	funk "github.com/thoas/go-funk"
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
			// create books if not exists
			for _, item := range request.Items {
				bookId, _, err := bookRepository.CreateIfNotExists(
					domain.Book{}.New(item.Book.toDataObject()),
					tx,
				)

				if err != nil {
					return nil, err
				}

				item.id = bookId
			}

			// create book receipt
			newReceipt := domain.BookReceipt{}.NewFromReceivingBooks(funk.Map(
				request.Items,
				func(item *ReceivingBook) *domain.ReceivingBook {
					return &domain.ReceivingBook{Book: item.toDataObject(), ReceivingQty: item.Qty}
				},
			).([]*domain.ReceivingBook))

			receiptId, err := bookReceiptRepository.Create(newReceipt, tx)
			if err != nil {
				return nil, err
			}

			return receiptId, err
		})

	return err
}
