package operation

import (
	"store/app/domain"
	tx "store/app/domain/transaction"
)

type receiveBooks struct{}

func (r *receiveBooks) Receive(request ReceiveBooksRequest) error {
	var receivingItems []domain.ReceivingBook
	for _, item := range request.Items {
		receivingItems = append(
			receivingItems,
			domain.ReceivingBook{Book: item.Book.toDataObject(), ReceivingQty: item.Qty},
		)
	}

	_, err := tx.ReceiveBooks(receivingItems)

	return err
}
