package command

import (
	"store/app/domain"
	repo "store/app/repository"

	funk "github.com/thoas/go-funk"
)

type ReceivingBook struct {
	Book
	Qty int
}

type ReceiveBooksRequest struct {
	Items []ReceivingBook
}

func (*command) Receive(request ReceiveBooksRequest) error {
	receivingBooks := map[string]*domain.ReceivingBook{}
	for _, item := range request.Items {
		receivingBooks[item.GoogleBookId] = &domain.ReceivingBook{
			Book:         item.Book.toDataObject(),
			ReceivingQty: item.Qty,
		}
	}

	_, err := TransactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			// create books if not exists
			for _, item := range request.Items {
				bookId, err := BookRepository.CreateIfNotExists(domain.Book{}.New(item.Book.toDataObject()), tx)
				if err != nil {
					return nil, err
				}

				receivingBooks[item.GoogleBookId].Id = bookId
			}

			// create book receipt
			newReceipt := domain.BookReceipt{}.NewFromReceivingBooks(funk.Map(
				receivingBooks,
				func(key string, value *domain.ReceivingBook) *domain.ReceivingBook {
					return value
				},
			).([]*domain.ReceivingBook))

			receiptId, err := BookReceiptRepository.Create(newReceipt, tx)
			if err != nil {
				return nil, err
			}

			return receiptId, err
		})

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

	orders, err := OrderRepository.GetReceivingOrders(nil)
	if err != nil {
		channel <- err
		return
	}

	for _, order := range orders {
		if ok, _ := order.UpdateToStockFilled(); ok {
			OrderRepository.Update(order, nil)
		}
	}
}
