package domain

import (
	"store/app/data"
	"store/app/domain"
)

func ReceiveBooks(receivingBooks []domain.ReceivingBook) (*domain.BookReceipt, error) {
	tx := domain.TransactionFactory.New()
	var err error

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// create book receipt
	receipt, err := domain.BookReceipt{}.Create(receivingBooks, tx)
	if err != nil {
		return nil, err
	}

	// increase onhand qty of receiving books
	stock := data.Stock{}
	for _, receivingBook := range receivingBooks {
		id, err := domain.Book{}.CreateIfNotExists(receivingBook.Book, tx)
		if err != nil {
			return nil, err
		}

		book, err := domain.Book{}.Get(id, tx)
		if err != nil {
			return nil, err
		}

		err = book.AdjustOnhandQty(receivingBook.ReceivingQty, tx)
		if err != nil {
			return nil, err
		}

		stock[id] = data.StockItem{
			BookId:       book.Id,
			OnhandQty:    book.OnhandQty + receivingBook.ReceivingQty,
			PreservedQty: book.PreservedQty,
		}
	}

	// Update order status to StockFilled for any orders
	// that can be fulfilled by the new stock
	orders, err := domain.Order{}.GetReceivingOrders(tx)
	if err == nil {
		for _, order := range orders {
			stock, err = order.TryUpdateToStockFilled(stock, tx)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return receipt, nil
}
