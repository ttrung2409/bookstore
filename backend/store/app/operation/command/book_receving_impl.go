package operation

import (
	"store/app/data"
	"store/app/domain"
)

type receiveBooks struct{}

func (*receiveBooks) Receive(request ReceiveBooksRequest) error {
	receivingBooks := []domain.ReceivingBook{}
	for _, item := range request.Items {
		receivingBooks = append(
			receivingBooks,
			domain.ReceivingBook{Book: item.toDataObject(), ReceivingQty: item.Qty},
		)
	}

	_, err := TransactionFactory.RunInTransaction(
		func(tx data.Transaction) (interface{}, error) {
			// create book receipt
			newReceipt := domain.BookReceipt{}.NewFromReceivingBooks(receivingBooks)
			receiptId, err := BookReceiptRepository.Create(newReceipt.State(), tx)
			if err != nil {
				return nil, err
			}

			// create books if not exists
			for _, item := range request.Items {
				dataBook := item.Book.toDataObject()
				_, err := BookRepository.CreateIfNotExists(&dataBook, tx)
				if err != nil {
					return nil, err
				}
			}

			dataReceipt, err := BookReceiptRepository.Get(receiptId, tx)
			if err != nil {
				return nil, err
			}

			receipt := domain.BookReceipt{}.New(dataReceipt)

			// increase on-hand qty of receiving books associated with the receipt items
			receipt.IncreaseStock()

			err = BookReceiptRepository.Update(receipt.State(), tx)

			return receipt, err
		}, nil)

	channel := make(chan error)

	go updateOrdersToStockFilled(channel)

	return err
}

// Update order status to StockFilled for any orders
// that can be fulfilled by the new stock
func updateOrdersToStockFilled(channel chan error) {
	defer func() {
		close(channel)
	}()

	dataOrders, err := OrderRepository.GetReceivingOrders(nil)
	if err != nil {
		channel <- err
		return
	}

	for _, dataOrder := range dataOrders {
		order := domain.Order{}.New(dataOrder)
		if ok, _ := order.TryUpdateToStockFilled(); ok {
			OrderRepository.Update(order.State(), nil)
		}
	}
}
