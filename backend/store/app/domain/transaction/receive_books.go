package domain

import (
	"store/app/data"
	"store/app/domain"
)

func ReceiveBooks(receivingBooks []domain.ReceivingBook) (*domain.BookReceipt, error) {
	receipt, err := domain.RunInTransaction(func(tx data.Transaction) (interface{}, error) {
		// create book receipt
		receipt, err := domain.BookReceipt{}.Create(receivingBooks, tx)
		if err != nil {
			return nil, err
		}

		// create books if not exists
		for _, receivingBook := range receivingBooks {
			_, err = domain.Book{}.CreateIfNotExists(receivingBook.Book, tx)
			if err != nil {
				return nil, err
			}
		}

		// increase on-hand qty of books associated with the receipt items
		stock, err := receipt.IncreaseStock(tx)
		if err != nil {
			return nil, err
		}

		// Update order status to StockFilled for any orders
		// that can be fulfilled by the new stock
		orders, err := domain.Order{}.GetReceivingOrders(tx)
		if err == nil {
			for _, order := range orders {
				stock, _ = order.TryUpdateToStockFilled(stock, tx)
			}
		}

		return receipt, nil
	}, nil)

	return receipt.(*domain.BookReceipt), err
}
