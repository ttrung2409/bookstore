package operation

import (
	"store/app/domain"
	transaction "store/app/domain/transaction"
)

type receiveBooks struct{}

func (r *receiveBooks) Receive(request ReceiveBooksRequest) error {
	var receivingItems []domain.ReceivingBook
	for _, item := range request.items {
		receivingItems = append(
			receivingItems,
			domain.ReceivingBook{Book: item.Book.toDataObject(), Qty: item.Qty},
		)
	}

	_, err := transaction.ReceiveBooks(receivingItems)

	return err
}
