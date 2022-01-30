package command

import (
	"store/app/domain"
	repo "store/app/repository"
	"store/container"
	"store/utils"

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
	var bookRepository = container.Instance().Get(utils.Nameof((*repo.BookRepository)(nil))).(repo.BookRepository)
	var bookReceiptRepository = container.Instance().Get(utils.Nameof((*repo.BookReceiptRepository)(nil))).(repo.BookReceiptRepository)
	var transactionFactory = container.Instance().Get(utils.Nameof((*repo.TransactionFactory)(nil))).(repo.TransactionFactory)

	receivingBooks := map[string]*domain.ReceivingBook{}
	for _, item := range request.Items {
		receivingBooks[item.GoogleBookId] = &domain.ReceivingBook{
			Book:         item.Book.toDataObject(),
			ReceivingQty: item.Qty,
		}
	}

	_, err := transactionFactory.RunInTransaction(
		func(tx repo.Transaction) (interface{}, error) {
			// create books if not exists
			for _, item := range request.Items {
				bookId, err := bookRepository.CreateIfNotExists(domain.Book{}.New(item.Book.toDataObject()), tx)
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

			receiptId, err := bookReceiptRepository.Create(newReceipt, tx)
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

	var orderRepository = container.Instance().Get(utils.Nameof((*repo.OrderRepository)(nil))).(repo.OrderRepository)

	orders, err := orderRepository.GetReceivingOrders(nil)
	if err != nil {
		channel <- err
		return
	}

	for _, order := range orders {
		if ok, _ := order.UpdateToStockFilled(); ok {
			orderRepository.Update(order, nil)
		}
	}
}
